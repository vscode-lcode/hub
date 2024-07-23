package hub

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/hashicorp/yamux"
	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Hub struct {
	locker  sync.RWMutex
	hosts   map[string]*Host
	hostTpl string
	domain  string
}

func New(domain, hostTpl string) *Hub {
	return &Hub{
		hosts:   map[string]*Host{},
		hostTpl: hostTpl,
		domain:  domain,
	}
}

var _ http.Handler = (*Hub)(nil)

func (hub *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostname := strings.SplitN(r.Host, ":", 2)[0]
	if strings.HasSuffix(hostname, hub.domain) {
		hub.locker.RLock()
		defer hub.locker.RUnlock()
		hhost, ok := hub.hosts[r.Host]
		if !ok || hhost == nil {
			http.Error(w, "host not found", http.StatusNotFound)
			return
		}
		hhost.ServeHTTP(w, r)
		return
	}
	hub.handleLink(w, r)
}

func (hub *Hub) handleLink(w http.ResponseWriter, r *http.Request) (err error) {
	defer err0.Then(&err, nil, nil)

	protocols := strings.SplitN(r.Header.Get("Sec-Websocket-Protocol"), ",", 4)
	if len(protocols) < 3 {
		http.Error(w, "can't find hostname and mac hash", http.StatusBadRequest)
		return
	}
	hostname, hash := protocols[1], protocols[2]
	hostname = strings.ToLower(hostname)
	host := fmt.Sprintf(hub.hostTpl, hostname)
	hhost := hub.SetHost(host, hash)
	if hhost == nil {
		http.Error(w, "init host failed", http.StatusUnauthorized)
		return
	}
	defer hub.freeHostIfNoSessions(hhost)

	h := w.Header()
	h.Set("X-Webdav-Host", host)

	socket := try.To1(websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols: []string{"webdav"},
	}))
	defer socket.Close(websocket.StatusAbnormalClosure, "")

	ctx := r.Context()

	var allow []string
	try.To(wsjson.Read(ctx, socket, &allow))

	conn := websocket.NetConn(ctx, socket, websocket.MessageBinary)

	yc := yamux.DefaultConfig()
	yc.LogOutput = os.Stdout
	sess := try.To1(yamux.Server(conn, yc))
	defer sess.Close()

	proxy := httputil.NewSingleHostReverseProxy(fakeTarget)
	proxy.Transport = &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return sess.Open()
		},
	}

	session := hhost.NewSession(proxy, allow)
	if session == nil {
		socket.Close(websocket.StatusAbnormalClosure, "init session failed")
		return
	}
	defer hhost.FreeSession(session)

	go func() {
		conn := try.To1(sess.Accept())
		defer conn.Close()
		hhost.flServer.ServeConn(conn)
	}()

	<-sess.CloseChan()
	return nil
}

var fakeTarget, _ = url.Parse("http://yamux.proxy")

func (hub *Hub) freeHostIfNoSessions(host *Host) {
	host.locker.RLock()
	defer host.locker.RUnlock()
	if len(host.sessions) > 0 {
		return
	}
	hub.locker.Lock()
	defer hub.locker.Unlock()
	delete(hub.hosts, host.Name)
}

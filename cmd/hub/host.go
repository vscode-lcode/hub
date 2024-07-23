package hub

import (
	"net/http"
	"net/rpc"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/vscode-lcode/hub/v2/cmd/hub/fl"
	"golang.org/x/net/webdav"
)

type Host struct {
	locker     sync.RWMutex
	Name       string
	Hash       string
	sessions   map[string]*Session
	fileLocker webdav.LockSystem
	flServer   *rpc.Server
}

func NewHost(name, hash string) *Host {
	ls := webdav.NewMemLS()

	srv := rpc.NewServer()
	fls := fl.New(ls)
	srv.Register(fls)

	host := &Host{
		Name:       name,
		Hash:       hash,
		sessions:   map[string]*Session{},
		fileLocker: ls,
		flServer:   srv,
	}
	return host
}

var _ http.Handler = (*Host)(nil)

func (hhost *Host) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := hhost.getAllowedSession(r.URL.Path)
	if sess == nil {
		http.Error(w, "no session allow this path", http.StatusForbidden)
		return
	}
	sess.ReverseProxy.ServeHTTP(w, r)
}

func (hhost *Host) getAllowedSession(path string) *Session {
	hhost.locker.RLock()
	defer hhost.locker.RUnlock()
	for _, sess := range hhost.sessions {
		for _, allow := range sess.Allow {
			if strings.HasPrefix(path, allow) {
				return sess
			}
		}
	}
	return nil
}

type Session struct {
	ID           string
	ReverseProxy http.Handler
	Allow        []string
}

func (hub *Hub) SetHost(name string, hash string) *Host {
	hub.locker.Lock()
	defer hub.locker.Unlock()
	host, ok := hub.hosts[name]
	if ok && host != nil {
		if host.Hash != hash {
			return nil
		}
		return host
	}
	host = NewHost(name, hash)
	hub.hosts[name] = host
	return host
}

func (hhost *Host) NewSession(rp http.Handler, allow []string) *Session {
	hhost.locker.Lock()
	defer hhost.locker.Unlock()
	sess := &Session{
		ReverseProxy: rp,
		Allow:        allow,
	}
	for range 3 {
		tid := uuid.NewString()
		_, ok := hhost.sessions[tid]
		if !ok {
			sess.ID = tid
			break
		}
	}
	if sess.ID == "" {
		return nil
	}
	hhost.sessions[sess.ID] = sess
	return sess
}

func (hhost *Host) FreeSession(sess *Session) {
	hhost.locker.Lock()
	defer hhost.locker.Unlock()

	delete(hhost.sessions, sess.ID)
}

package main

import (
	"fmt"
	"net/http"
	"strings"

	"gitlab.com/jonas.jasas/httprelay/pkg/controller"
	"gitlab.com/jonas.jasas/httprelay/pkg/repository"
)

func NewProxy(stopChan chan struct{}) http.Handler {
	proxyRep := repository.NewProxyRep()
	proxyCtrl := controller.NewProxyCtrl(proxyRep, stopChan)
	fn := wildcardAgentExistsHandler(proxyCtrl, proxyCtrl.Conduct)
	return wildcardCorsHandler(fn)
}

func wildcardAgentExistsHandler(pc *controller.ProxyCtrl, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.EqualFold(r.Method, "SERVE") && !pc.HasProxySer(r.URL.Path) {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, "agent is not ready")
			return
		}
		h(w, r)
	}
}

func wildcardCorsHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		wildcardCors(w, r)
		if r.Method != "OPTIONS" {
			h(w, r)
		}
	}
}

const WebdavExtraMethods = ", PROPFIND, COPY, MKCOL, LOCK, UNLOCK, MOVE"

func wildcardCors(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin == "" {
		origin = "*"
	}
	w.Header().Set("Access-Control-Allow-Origin", origin)
	//w.Header().Set("HttpRelay-Version", Version)

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, PATCH, DELETE, OPTIONS, SERVE"+WebdavExtraMethods)
		w.Header().Set("Access-Control-Allow-Headers", "*")
	} else {
		w.Header().Set("Access-Control-Expose-Headers", "*")
	}
}

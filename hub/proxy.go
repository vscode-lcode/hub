package main

import (
	"net/http"

	"gitlab.com/jonas.jasas/httprelay/pkg/controller"
	"gitlab.com/jonas.jasas/httprelay/pkg/repository"
)

func NewProxy(stopChan chan struct{}) http.Handler {
	proxyRep := repository.NewProxyRep()
	proxyCtrl := controller.NewProxyCtrl(proxyRep, stopChan)
	return wildcardCorsHandler(proxyCtrl.Conduct)
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

func wildcardCors(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin == "" {
		origin = "*"
	}
	w.Header().Set("Access-Control-Allow-Origin", origin)
	//w.Header().Set("HttpRelay-Version", Version)

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, PATCH, DELETE, OPTIONS, SERVE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
	} else {
		w.Header().Set("Access-Control-Expose-Headers", "*")
	}
}

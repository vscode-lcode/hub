package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/donovanhide/eventsource"
)

type OpenEvent string

func (t OpenEvent) Id() string    { return fmt.Sprint(time.Now().UnixNano()) }
func (t OpenEvent) Event() string { return "open-webdav" }
func (t OpenEvent) Data() string  { return string(t) }

var _ eventsource.Event = OpenEvent("")

func initOpener(mux *http.ServeMux) sync.WaitGroup {
	srv := eventsource.NewServer()
	var wg sync.WaitGroup

	var openerSSEHandler = srv.Handler("open")
	http.HandleFunc("/open-handler", func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1)
		defer wg.Done()
		openerSSEHandler(w, r)
	})

	http.HandleFunc("/open", func(w http.ResponseWriter, r *http.Request) {
		link := r.URL.Query().Get("link")
		if link == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "lost link query field")
			return
		}
		srv.Publish([]string{"open"}, OpenEvent(link))
		fmt.Fprint(w, "open event has been sent")
	})

	return wg
}

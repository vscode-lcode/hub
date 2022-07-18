package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	parseFlag()

	if had, err := hasRunningHub(); had && err == nil {
		return
	}

	stopChan := make(chan struct{})
	mux := http.DefaultServeMux

	mux.Handle("/proxy/", NewProxy(stopChan))
	var wg = initOpener(mux)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, healthText)
	})

	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		if args.Foreground {
			fmt.Fprint(w, "exit signal was ignored")
			return
		}
		forceExit := r.URL.Query().Has("force")
		go exitProgram(wg, forceExit)
		fmt.Fprint(w, "exit signal has received")
	})
	http.HandleFunc("/open-link", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		p := q.Get("path")
		if p == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "lost path field")
			return
		}
		link := GetOpenLink(p, q.Get("link"))
		fmt.Fprint(w, link)
	})

	log.Fatal(http.ListenAndServe(args.Addr, mux))

}

func exitProgram(wg *sync.WaitGroup, force bool) {
	if force {
		os.Exit(0)
		return
	}
	wg.Wait()
	os.Exit(0)
}

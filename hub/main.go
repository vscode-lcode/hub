package main

import (
	"flag"
	"fmt"
	"net/http"
)

var args struct {
	Addr string
}

func init() {
	flag.StringVar(&args.Addr, "l", "127.0.0.1:4349", "Bind address")
	flag.Parse()
}

func main() {
	stopChan := make(chan struct{})

	http.Handle("/proxy/", NewProxy(stopChan))
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "it is working")
	})
	http.HandleFunc("/open", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "open")
	})

	http.ListenAndServe(args.Addr, nil)
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
	mux := http.DefaultServeMux

	mux.Handle("/proxy/", NewProxy(stopChan))
	var wg = initOpener(mux)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "it is working")
	})
	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		wg.Wait()
		os.Exit(0)
	})

	log.Fatal(http.ListenAndServe(args.Addr, mux))

}

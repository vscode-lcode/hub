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

	var exitProgram = func() {
		wg.Wait()
		os.Exit(0)
	}
	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		go exitProgram()
		fmt.Fprint(w, "exit signal has received")
	})

	log.Fatal(http.ListenAndServe(args.Addr, mux))

}

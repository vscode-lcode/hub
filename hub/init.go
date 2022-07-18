package main

import (
	"flag"
	"fmt"
	"net/url"
)

var args struct {
	Addr           string
	Foreground     bool
	LinkEditorType string
	PWACode        string
}

func init() {
	flag.StringVar(&args.Addr, "l", "127.0.0.1:4349", "Bind address")
	flag.BoolVar(&args.Foreground, "f", false, "Run Hub foreground")
	flag.StringVar(&args.LinkEditorType, "link", "vscode", "editor type is vscode or browser")
	flag.StringVar(&args.PWACode, "pwa", "https://vscode.dev", "the site of vscode web")
}

var (
	apiHealthAddr  string
	apiReqExitAddr string
	PWACode        *url.URL
)

func parseFlag() {
	flag.Parse()

	var err error
	PWACode, err = url.Parse(args.PWACode)
	if err != nil {
		panic(err)
	}
	apiHealthAddr = fmt.Sprintf("http://%s/health", args.Addr)
	apiReqExitAddr = fmt.Sprintf("http://%s/exit", args.Addr)
}

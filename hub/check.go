package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var healthText = fmt.Sprintf("it is working. v%s", VERSION)

func hasRunningHub() (running bool, err error) {
	resp, err := http.Get(apiHealthAddr)
	if err != nil {
		return
	}
	ht, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	isSameVersion := string(ht) == healthText
	if !isSameVersion {
		defer time.Sleep(500 * time.Millisecond)
		if _, err = http.Get(apiReqExitAddr + "?force"); err != nil {
			return
		}
		return
	}
	running = true
	return
}

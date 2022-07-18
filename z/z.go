package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.Parse("http://127.0.0.1:447")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v \n", u)
}

package main

import (
	"fmt"
	"net/http"
)

type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter=%d", ctr.n)
}

func ootest() {
	ctr := new(Counter)
	http.Handle("/counter", ctr)
}


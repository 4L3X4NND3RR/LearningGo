package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

// UUID is a custom multiplexer
type UUID struct {
}

func (p *UUID) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.URL.Path == "/" {
		giveRandomUUID(rw, rq)
		return
	}
	http.NotFound(rw, rq)
	return
}

func giveRandomUUID(rw http.ResponseWriter, rq *http.Request) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(rw, fmt.Sprintf("%x", b))
}

func main() {
	mux := &UUID{}
	http.ListenAndServe(":8000", mux)
}

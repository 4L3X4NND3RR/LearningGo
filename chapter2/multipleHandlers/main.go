package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/randomFloat", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, rand.Float64())
	})
	newMux.HandleFunc("/randomInt", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, rand.Intn(100))
	})
	http.ListenAndServe(":8000", newMux)
}

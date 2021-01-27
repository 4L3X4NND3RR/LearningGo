package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func QueryHandler(rw http.ResponseWriter, rq *http.Request)  {
	queryParams := rq.URL.Query()
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Got parameter id: %s!\n", queryParams["id"][0])
	fmt.Fprintf(rw, "Got parameter category: %s!\n", queryParams["category"][0])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles", QueryHandler)
	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

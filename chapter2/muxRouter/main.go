package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ArticleHandler(rw http.ResponseWriter, rq *http.Request) {
	vars := mux.Vars(rq)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Category is: %v\n", vars["category"])
	fmt.Fprintf(rw, "ID is: %v\n", vars["id"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

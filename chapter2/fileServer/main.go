package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("/Users/4L3X4NND3RR/static"))
	log.Fatal(http.ListenAndServe(":8000", router))
}

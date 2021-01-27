package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func HealthCheck(rw http.ResponseWriter, req *http.Request)  {
	currentTime := time.Now()
	io.WriteString(rw, currentTime.String())
}

func main() {
	http.HandleFunc("/health", HealthCheck)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

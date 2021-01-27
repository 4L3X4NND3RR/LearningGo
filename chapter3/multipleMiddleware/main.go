package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area uint64
}

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")
		// Filtering requests by MIME type
		if r.Header.Get("Content-Type") != "application/json" {
			rw.WriteHeader(http.StatusUnsupportedMediaType)
			rw.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
			return
		}
		handler.ServeHTTP(rw, r)
	})
}

func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(rw, r)
		// Setting cookie to every API response
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(rw, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}

func handle(rw http.ResponseWriter, r *http.Request) {
		// Check if method is POST
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		// Your resource creation logic goes here. For now it is plain print to console
		log.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
		// Tell everything is fine
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("201 - Created"))
	} else {
		// Say method not allowed
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte("405 - Method Not Allowed"))
	}
}

func main() {
	originalHandler := http.HandlerFunc(handle)
	http.Handle("/city", filterContentType(setServerTimeCookie(originalHandler)))
	http.ListenAndServe(":8000", nil)
}

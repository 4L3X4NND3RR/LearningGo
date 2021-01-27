package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type city struct {
	Name string
	Area uint64
}

func postHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)	
		}
		defer r.Body.Close()
		fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("201 - Created"))
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte("405 - Method Not Allowed"))
	}
}

func main() {
	http.HandleFunc("/city", postHandler)
	http.ListenAndServe(":8000", nil)
}

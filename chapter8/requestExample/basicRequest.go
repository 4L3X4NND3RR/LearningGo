package main

import (
	"log"

	"github.com/levigross/grequests"
)

func main() {
	res, err := grequests.Get("http://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	log.Println(res.String())
}

package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

func pingTime(rq *restful.Request, rp *restful.Response) {
	io.WriteString(rp, fmt.Sprintf("%s", time.Now()))
}

func main() {
	// create a web service
	webservice := new(restful.WebService)
	// Create a route and attach it to handler in the service
	webservice.Route(webservice.GET("/ping").To(pingTime))
	// add the service to the application
	restful.Add(webservice)
	http.ListenAndServe(":8000", nil)
}

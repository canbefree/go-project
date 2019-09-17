package main

import (
	"net/http"

	"github.com/emicklei/go-restful"
	// "github.com/emicklei/go-restful-swagger12"
)

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/").
		Doc("get all users").
		Writes("").
		Returns(200, "OK", "ok"))

	container := restful.NewContainer().Add(ws)
	http.ListenAndServe(":8080", container)
}

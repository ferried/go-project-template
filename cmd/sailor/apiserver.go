package main

import (
	"log"
	"net/http"
	"sailor/pkg/apis"

	"github.com/emicklei/go-restful"
)

func main() {
	u := apis.UserResource{}
	restful.DefaultContainer.Add(u.WebService())
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/ferried/Projects/sailor/docs"))))
	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package apis

import "github.com/emicklei/go-restful"

type User struct {
	ID   string `json:"id" description:"identifier of the user"`
	Name string `json:"name" description:"name of the user" default:"john"`
	Age  int    `json:"age" description:"age of the user" default:"21"`
}

func Users(request *restful.Request, response *restful.Response) {
	list := []User{User{"id", "a", 13}}
	response.WriteAsJson(list)
}

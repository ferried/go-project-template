package apis

import (
	"github.com/emicklei/go-restful"
	apis "sailor/pkg/apis/users"
)

type UserResource struct {
	users map[string]apis.User
}

func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/users").To(apis.Users))
	return ws;
}

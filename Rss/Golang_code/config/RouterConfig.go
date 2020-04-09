package config

import (

	"RSs/handle/userHandle"
	"RSs/plugins/auth"
	"github.com/emicklei/go-restful"
)

var user = new(userHandle.User)


const (
	PATH = "/api"
)

// restful容器初始化
var wc = restful.NewContainer()

// restful 服务初始化
var ws = new(restful.WebService)

func routerConf() {
	usersRouterConf()

}

func GetRouterContainer() *restful.Container {

	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path(PATH)
	routerConf()
	wc.Add(ws)
	return wc
}

func usersRouterConf() {
	rootPath := "users/"
	methodPath := "login"
	ws.Route(ws.POST(rootPath + methodPath).To(user.Login))

	methodPath = "addUser"
	ws.Route(ws.POST(rootPath + methodPath).To(user.AddUser))
	methodPath = "get_info"
	ws.Route(ws.GET(rootPath + methodPath).To(user.GetInfo))

	methodPath = "logout"
	ws.Route(ws.GET(rootPath + methodPath).To(user.Logout).Filter(auth.TokenFilter))
}
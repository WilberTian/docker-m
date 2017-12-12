package routers

import (
	"docker-m/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    /* HTTP API for docker remote API */
	beego.Router("/dockerapi/containers/json", &controllers.ContainerController{}, "get:GetContainers")
	beego.Router("/dockerapi/containers/:id/json", &controllers.ContainerController{}, "get:GetContainer")
	beego.Router("/dockerapi/containers/:id/top", &controllers.ContainerController{}, "get:TopContainer")
	beego.Router("/dockerapi/containers/:id/start", &controllers.ContainerController{}, "post:StartContainer")
	beego.Router("/dockerapi/containers/:id/stop", &controllers.ContainerController{}, "post:StopContainer")
	beego.Router("/dockerapi/containers/:id", &controllers.ContainerController{}, "delete:DeleteContainer")
	beego.Router("/dockerapi/containers/:id/stats", &controllers.ContainerController{}, "get:GetContainerStats")


	beego.Router("/dockerapi/images/json", &controllers.ImageController{}, "get:GetImages")
	beego.Router("/dockerapi/images/:id/json", &controllers.ImageController{}, "get:GetImage")
	beego.Router("/dockerapi/images/:user/:repo/json", &controllers.ImageController{}, "get:GetUserImage")
	beego.Router("/dockerapi/images/:id", &controllers.ImageController{}, "delete:DeleteImage")

	
	beego.Router("/dockerapi/version", &controllers.MiscController{}, "get:GetVersion")
	beego.Router("/dockerapi/info", &controllers.MiscController{}, "get:GetInfo")
	beego.Router("/dockerapi/images/search", &controllers.MiscController{}, "get:GetSearchImages")
	// beego.Router("/dockerapi/events", &controllers.DockerapiController{}, "get:GetEvents") // Not support yet
	beego.Router("/dockerapi/_ping", &controllers.MiscController{}, "get:Ping")
}

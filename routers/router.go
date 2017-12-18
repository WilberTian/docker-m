package routers

import (
	"docker-m/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/docker-m/validateLogin", &controllers.LoginController{}, "get:ValidateLogin")

	beego.Router("/docker-m/containers", &controllers.ContainerController{}, "get:GetContainers")
	beego.Router("/docker-m/containers/:id", &controllers.ContainerController{}, "get:GetContainerDetail")
	beego.Router("/docker-m/containers/:id/top", &controllers.ContainerController{}, "get:GetContainerProcesses")
	beego.Router("/docker-m/containers/:id/start", &controllers.ContainerController{}, "post:StartContainer")
	beego.Router("/docker-m/containers/:id/stop", &controllers.ContainerController{}, "post:StopContainer")
	beego.Router("/docker-m/containers/:id/delete", &controllers.ContainerController{}, "delete:DeleteContainer")
	beego.Router("/docker-m/containers/:id/stats", &controllers.ContainerController{}, "get:GetContainerStats")

	beego.Router("/docker-m/images", &controllers.ImageController{}, "get:GetImages")
	/*
		beego.Router("/docker-m/images/:id", &controllers.ImageController{}, "get:GetImage")
		beego.Router("/docker-m/images/:user/:repo", &controllers.ImageController{}, "get:GetUserImage")
		beego.Router("/docker-m/images/:id", &controllers.ImageController{}, "delete:DeleteImage")
	*/
	/*
		beego.Router("/docker-m/version", &controllers.MiscController{}, "get:GetVersion")
		beego.Router("/docker-m/info", &controllers.MiscController{}, "get:GetInfo")
		beego.Router("/docker-m/images/search", &controllers.MiscController{}, "get:GetSearchImages")
		// beego.Router("/docker-m/events", &controllers.DockerapiController{}, "get:GetEvents") // Not support yet
		beego.Router("/docker-m/_ping", &controllers.MiscController{}, "get:Ping")
	*/
}

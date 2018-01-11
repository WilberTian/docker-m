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
	beego.Router("/docker-m/containers/:id/delete", &controllers.ContainerController{}, "post:DeleteContainer")
	beego.Router("/docker-m/containers/:id/stats", &controllers.ContainerController{}, "get:GetContainerStats")

	beego.Router("/docker-m/images", &controllers.ImageController{}, "get:GetImages")
	beego.Router("/docker-m/images/:id", &controllers.ImageController{}, "get:GetImageDetail")
	beego.Router("/docker-m/images/:id/history", &controllers.ImageController{}, "get:GetImageHistory")
	beego.Router("/docker-m/images/:id/delete", &controllers.ImageController{}, "post:DeleteImage")

	beego.Router("/docker-m/version", &controllers.MiscController{}, "get:GetVersion")
	beego.Router("/docker-m/info", &controllers.MiscController{}, "get:GetInfo")
	beego.Router("/docker-m/_ping", &controllers.MiscController{}, "get:Ping")
}

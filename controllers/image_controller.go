package controllers

import (
	"github.com/astaxie/beego"

	"docker-m/utils"
)

type ImageController struct {
	beego.Controller
}


func (this *ImageController) GetImages() {
	address := "/images/json"
	result := utils.InitDockerConnection(address, "GET")
	this.Ctx.WriteString(result)
}

func (this *ImageController) GetImage() {
	id := this.GetString(":id")
	address := "/images/" + id + "/json"
	result := utils.InitDockerConnection(address, "GET")
	this.Ctx.WriteString(result)
}

func (this *ImageController) GetUserImage() {
	user := this.GetString(":user")
	repo := this.GetString(":repo")
	address := "/images/" + user + "/" + repo + "/json"
	result := utils.InitDockerConnection(address, "GET")
	this.Ctx.WriteString(result)
}

func (this *ImageController) DeleteImage() {
	id := this.GetString(":id")
	address := "/images/" + id
	result := utils.InitDockerConnection(address, "DELETE")
	this.Ctx.WriteString(result)
}

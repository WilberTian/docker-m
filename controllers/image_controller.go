package controllers

import (
	"github.com/astaxie/beego"

	"docker-m/vos"

	"docker-m/services"
)

type ImageController struct {
	beego.Controller
}

func (this *ImageController) GetImages() {
	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	data, err := services.GetImages()
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	} else {
		responseVo.Data = data
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

/*
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
*/

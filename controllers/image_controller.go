package controllers

import (
	"github.com/astaxie/beego"

	"docker-m/vos"

	"docker-m/services"

	"docker-m/utils"
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


func (this *ImageController) GetImageDetail() {
	id := this.GetString(":id")

	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	data, err := services.GetImageDetail(id)
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	} else {
		responseVo.Data = data
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

func (this *ImageController) GetImageHistory() {
	id := this.GetString(":id")

	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	data, err := services.GetImageHistory(id)
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	} else {
		responseVo.Data = data
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}


func (this *ImageController) DeleteImage() {
	id := this.GetString(":id")

	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	params := utils.GetQueryString(this.Ctx.Input.URI())
	err := services.DeleteImage(id, params)
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

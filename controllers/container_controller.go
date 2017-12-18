package controllers

import (
	"github.com/astaxie/beego"

	"docker-m/vos"

	"docker-m/utils"

	"docker-m/services"
)

type ContainerController struct {
	beego.Controller
}

func (this *ContainerController) GetContainers() {
	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	params := utils.GetQueryString(this.Ctx.Input.URI())
	data, err := services.GetContainers(params)
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	} else {
		responseVo.Data = data
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

func (this *ContainerController) GetContainerDetail() {
	id := this.GetString(":id")

	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	data, err := services.GetContainerDetail(id)
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	} else {
		responseVo.Data = data
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

func (this *ContainerController) GetContainerProcesses() {
	id := this.GetString(":id")

	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	data, err := services.GetContainerProcesses(id)
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	} else {
		responseVo.Data = data
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

func (this *ContainerController) StartContainer() {
	id := this.GetString(":id")

	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	err := services.StartContainer(id)
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

func (this *ContainerController) StopContainer() {
	id := this.GetString(":id")

	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	err := services.StopContainer(id)
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

func (this *ContainerController) DeleteContainer() {
	id := this.GetString(":id")

	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	params := utils.GetQueryString(this.Ctx.Input.URI())
	err := services.DeleteContainer(id, params)
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

func (this *ContainerController) GetContainerStats() {
	id := this.GetString(":id")

	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	params := utils.GetQueryString(this.Ctx.Input.URI())
	data, err := services.GetContainerStats(id, params)
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	} else {
		responseVo.Data = data
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

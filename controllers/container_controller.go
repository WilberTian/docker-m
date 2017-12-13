package controllers

import (
	"github.com/astaxie/beego"

	"docker-m/utils"
)

type ContainerController struct {
	beego.Controller
}


func (this *ContainerController) GetContainers() {
	address := "/containers/json"
	address = address + "?" + utils.GetQueryString(this.Ctx.Input.URI())
	result := utils.InitDockerConnection(address, "GET")
	this.Ctx.WriteString(result)
}

func (this *ContainerController) GetContainer() {
	id := this.GetString(":id")
	address := "/containers/" + id + "/json"
	result := utils.InitDockerConnection(address, "GET")
	this.Ctx.WriteString(result)
}

func (this *ContainerController) TopContainer() {
	id := this.GetString(":id")
	address := "/containers/" + id + "/top"
	result := utils.InitDockerConnection(address, "GET")
	this.Ctx.WriteString(result)
}

func (this *ContainerController) StartContainer() {
	id := this.GetString(":id")
	address := "/containers/" + id + "/start"
	result := utils.InitDockerConnection(address, "POST")
	this.Ctx.WriteString(result)
}

func (this *ContainerController) StopContainer() {
	id := this.GetString(":id")
	address := "/containers/" + id + "/stop"
	result := utils.InitDockerConnection(address, "POST")
	this.Ctx.WriteString(result)
}

func (this *ContainerController) DeleteContainer() {
	id := this.GetString(":id")
	address := "/containers/" + id
	result := utils.InitDockerConnection(address, "DELETE")
	this.Ctx.WriteString(result)
}

func (this *ContainerController) GetContainerStats() {
	id := this.GetString(":id")
	address := "/containers/" + id + "/stats?stream=False"
	result := utils.InitDockerConnection(address, "GET")
	this.Ctx.WriteString(result)
}

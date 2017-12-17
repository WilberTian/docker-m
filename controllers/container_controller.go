package controllers

import (
	"encoding/json"

	"time"

	"github.com/astaxie/beego"

	"docker-m/utils"

	"docker-m/vos"
)

type ContainerController struct {
	beego.Controller
}

type PortVo struct {
	PrivatePort int64
	PublicPort  int64
	Type        string
	IP          string
}

type ContainersVo struct {
	ID      string `json:"Id"`
	Image   string
	Command string
	Created int64
	State   string
	Status  string
	Ports   []PortVo
	Names   []string
}

type ContainerNetworkVo struct {
	Aliases             []string
	MacAddress          string
	GlobalIPv6PrefixLen int
	GlobalIPv6Address   string
	IPv6Gateway         string
	IPPrefixLen         int
	IPAddress           string
	Gateway             string
	EndpointID          string
	NetworkID           string
}

type PortBindingVo struct {
	HostIP   string
	HostPort string
}

type NetworkSettingsVo struct {
	Networks               map[string]ContainerNetworkVo
	IPAddress              string
	IPPrefixLen            int
	MacAddress             string
	Gateway                string
	Bridge                 string
	Ports                  map[PortVo][]PortBindingVo
	NetworkID              string
	EndpointID             string
	SandboxKey             string
	GlobalIPv6Address      string
	GlobalIPv6PrefixLen    int
	IPv6Gateway            string
	LinkLocalIPv6Address   string
	LinkLocalIPv6PrefixLen int
	SecondaryIPAddresses   []string
	SecondaryIPv6Addresses []string
}


type ContainerDetailVo struct {
	ID      		string `json:"Id"`
	Created 		time.Time
	Path    		string
	Args    		[]string
	Image   		string
	NetworkSettings *NetworkSettingsVo
	Name            string
	Driver          string
	Volumes         map[string]string
}

func (this *ContainerController) GetContainers() {
	address := "/containers/json"
	address = address + "?" + utils.GetQueryString(this.Ctx.Input.URI())
	result := utils.InitDockerConnection(address, "GET")

	var containers []ContainersVo
	json.Unmarshal([]byte(result), &containers)

	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    containers,
		Success: true,
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

func (this *ContainerController) GetContainer() {
	id := this.GetString(":id")
	address := "/containers/" + id + "/json"
	result := utils.InitDockerConnection(address, "GET")

	var containerDetail ContainerDetailVo
	
	json.Unmarshal([]byte(result), &containerDetail)

	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    containerDetail,
		Success: true,
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
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

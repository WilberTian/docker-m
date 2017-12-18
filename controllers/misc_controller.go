package controllers

/*
import (
	"github.com/astaxie/beego"

	"docker-m/utils"
)

type MiscController struct {
	beego.Controller
}

func (this *MiscController) GetVersion() {
	address := "/version"
	result := utils.InitDockerConnection(address, "GET")
	this.Ctx.WriteString(result)
}

func (this *MiscController) GetInfo() {
	address := "/info"
	result := utils.InitDockerConnection(address, "GET")
	this.Ctx.WriteString(result)
}

func (this *MiscController) GetSearchImages() {
	address := "/images/search"
	address = address + "?" + utils.GetQueryString(this.Ctx.Input.URI())
	result := utils.InitDockerConnection(address, "GET")
	this.Ctx.WriteString(result)
}
*/
/* Todo: Implement events API, the response is a stream so can't use ioutil.ReadAll() which will be blocked
func (this *MiscController) GetEvents() {
	address := "/events"
	var since int
	this.Ctx.Input.Bind(&since, "since")
	address = address + "?since=" + strconv.Itoa(since)
	result := utils.InitDockerConnection(address, "GET")
	this.Ctx.WriteString(result)
}
*/
/*
func (this *MiscController) Ping() {
	address := "/_ping"
	result := utils.InitDockerConnection(address, "GET")
	this.Ctx.WriteString(result)
}
*/

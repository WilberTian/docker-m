package controllers


import (
	"github.com/astaxie/beego"

	"docker-m/vos"

	"docker-m/services"
)

type MiscController struct {
	beego.Controller
}

func (this *MiscController) GetVersion() {
	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	data, err := services.GetVersion()
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	} else {
		responseVo.Data = data
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}


func (this *MiscController) GetInfo() {
	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	data, err := services.GetInfo()
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	} else {
		responseVo.Data = data
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}


func (this *MiscController) Ping() {
	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    "",
		Success: true,
	}

	err := services.Ping()
	if err != nil {
		responseVo.Code = 500
		responseVo.Message = err.Error()
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

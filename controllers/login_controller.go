package controllers

import (
	"github.com/astaxie/beego"

	"docker-m/vos"
)

type LoginVo struct {
	Username string `json:"username"`
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) ValidateLogin() {
	loginVo := LoginVo{
		Username: "wilber",
	}

	responseVo := vos.ResponseVo{
		Code:    200,
		Message: "",
		Data:    loginVo,
		Success: true,
	}

	this.Data["json"] = &responseVo
	this.ServeJSON()
}

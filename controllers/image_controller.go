package controllers

import (
	"encoding/json"
	
	"github.com/astaxie/beego"

	"docker-m/utils"

	"docker-m/vos"
)

type ImageController struct {
	beego.Controller
}

type ImagesVo struct {
	ID string `json:"Id"`
	RepoTags []string          
	Created int64             
	Size int64             
	VirtualSize int64             
	ParentID string            
	RepoDigests []string          
	Labels map[string]string 
}

func (this *ImageController) GetImages() {
	address := "/images/json"
	result := utils.InitDockerConnection(address, "GET")
	
	var iamges []ImagesVo
	json.Unmarshal([]byte(result), &iamges)

	responseVo := vos.ResponseVo {
		Code: 200,
		Message: "",
		Data: iamges,
	    Success: true,
	}

	this.Data["json"] = &responseVo
    this.ServeJSON()
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

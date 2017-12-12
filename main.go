package main

import (
	_ "docker-m/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}


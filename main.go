package main

import (
	"github.com/astaxie/beego"
	_ "gogs.mebox.wiki/course-plus/courseplus-back-end/filters"
	_ "gogs.mebox.wiki/course-plus/courseplus-back-end/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"plivo/plivo.persistence/postgre"
	"plivo/plivo.persistence/redis"
	_ "plivo/plivo.sms.api/routers"

	"github.com/astaxie/beego"
)

func main() {
	if postgre.Initialize() != nil {
		fmt.Println("Failed to inititalize Postgre. Exiting...")

		return
	}

	redis.Initialize()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}

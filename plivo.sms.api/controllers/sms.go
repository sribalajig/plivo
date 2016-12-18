package controllers

import (
	"github.com/astaxie/beego"
)

type SMSController struct {
	beego.Controller
}

func (smsController *SMSController) InBound() {
	smsController.Data["json"] = "SMS controller works"

	smsController.ServeJSON()
}

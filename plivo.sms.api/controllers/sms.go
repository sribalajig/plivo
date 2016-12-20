package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"plivo/plivo.sms/model"
	"plivo/plivo.sms/service"
	"plivo/plivo.sms/validator"
)

type SMSController struct {
	beego.Controller
}

func (smsController *SMSController) Inbound() {
	var sms model.SMS

	json.Unmarshal(smsController.Ctx.Input.RequestBody, &sms)

	err := service.NewSMSService().ProcessInboundSMS(sms)

	if err == nil {
		smsController.Ctx.Output.JSON(validator.ValidationResult{Message: "inbound sms ok”"}, false, false)
	} else {
		smsController.Ctx.Output.SetStatus(500)
		smsController.Ctx.Output.JSON(validator.ValidationResult{Error: err.Error()}, false, false)
	}
}

func (smsController *SMSController) Outbound() {
	var sms model.SMS

	json.Unmarshal(smsController.Ctx.Input.RequestBody, &sms)

	err := service.NewSMSService().ProcessOutboundSMS(sms)

	if err == nil {
		smsController.Ctx.Output.JSON(validator.ValidationResult{Message: "outbound sms ok"}, false, false)
	} else {
		smsController.Ctx.Output.SetStatus(500)
		smsController.Ctx.Output.JSON(validator.ValidationResult{Error: err.Error()}, false, false)
	}

	smsController.ServeJSON()
}

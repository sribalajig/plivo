package filters

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"plivo/plivo.sms/model"
	"plivo/plivo.sms/validator"
)

var FilterSMS = func(ctx *context.Context) {
	var sms model.SMS

	json.Unmarshal(ctx.Input.RequestBody, &sms)

	if validationResult := validator.NewSMSValidator().Validate(sms); !validationResult.IsSuccess {
		ctx.Output.SetStatus(400)
		ctx.Output.JSON(validationResult, false, true)
	}
}

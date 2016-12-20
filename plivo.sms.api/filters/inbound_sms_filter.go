package filters

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"plivo/plivo.core/phone_numbers/repository"
	"plivo/plivo.sms/model"
	"plivo/plivo.sms/validator"
)

var FilterInboundSMS = func(ctx *context.Context) {
	var sms model.SMS

	json.Unmarshal(ctx.Input.RequestBody, &sms)

	numberExists := repository.PhoneNumberRepository{}.Exists(sms.To)

	if !numberExists {
		ctx.Output.SetStatus(400)
		ctx.Output.JSON(validator.ValidationResult{
			Error: "to parameter not found",
		}, false, true)
	}

}

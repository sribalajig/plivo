package filters

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"plivo/plivo.core/phone_numbers/repository"
	"plivo/plivo.sms/model"
	"plivo/plivo.sms/validator"
)

var FilterOutboundSMS = func(ctx *context.Context) {
	var sms model.SMS

	json.Unmarshal(ctx.Input.RequestBody, &sms)

	numberExists := repository.PhoneNumberRepository{}.Exists(sms.From)

	if !numberExists {
		ctx.Output.SetStatus(400)
		ctx.Output.JSON(validator.ValidationResult{
			Error: "from parameter not found",
		}, false, true)
	}

}

package filters

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"plivo/plivo.core/phone_numbers/repository"
	"plivo/plivo.sms/model"
	"plivo/plivo.sms/service"
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

	smsService := service.NewSMSService()

	if isStopped, _ := smsService.IsStopped(sms); isStopped {
		ctx.Output.SetStatus(403)

		ctx.Output.JSON(validator.ValidationResult{
			Error: fmt.Sprintf("sms from %s to %s blocked by STOP request", sms.From, sms.To),
		}, false, true)
	}

	if count, _ := smsService.GetSentSMSCount(sms); count > 50 {
		ctx.Output.SetStatus(403)

		ctx.Output.JSON(validator.ValidationResult{
			Error: fmt.Sprintf("limit reached for from %s", sms.From),
		}, false, true)
	}
}

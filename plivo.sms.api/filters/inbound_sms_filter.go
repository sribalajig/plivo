package filters

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"plivo/plivo.sms/model"
)

var FilterInboundSMS = func(ctx *context.Context) {
	var inboundSMS model.InboundSMS

	json.Unmarshal(ctx.Input.RequestBody, &inboundSMS)

	if !inboundSMS.Validate() {
		ctx.Output.SetStatus(400)
		ctx.Output.JSON(&struct{ message string }{"There was an error in parsing the request"}, false, true)
	}
}

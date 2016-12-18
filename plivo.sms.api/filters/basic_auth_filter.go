package filters

import (
	"github.com/astaxie/beego/context"
	"plivo/plivo.core/authentication"
	"plivo/plivo.sms/validator"
)

var BasicAuthFilter = func(ctx *context.Context) {
	authValidationResult := validator.ValidationResult{
		Error: "You are not authorized to view the requested information",
	}

	success, creds := authentication.GetCredentials(ctx.Request.Header.Get("Authorization"))

	if !success {
		ctx.Output.SetStatus(403)
		ctx.Output.JSON(authValidationResult, false, true)
		return
	}

	authorized := authentication.NewBasicAuthenticator().Authenticate(*creds)

	if !authorized {
		ctx.Output.SetStatus(403)
		ctx.Output.JSON(authValidationResult, false, true)
	}
}

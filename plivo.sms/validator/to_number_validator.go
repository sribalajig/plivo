package validator

import (
	"plivo/plivo.sms/model"
)

type ToNumberValidator struct {
}

func (toNumberValidator ToNumberValidator) Validate(sms model.SMS) ValidationResult {
	if len(sms.To) == 0 {
		return ValidationResult{
			Message: "to is missing",
		}
	}

	if !validate(sms.To) {
		return ValidationResult{
			Message: "to is invalid",
		}
	}

	return ValidationResult{
		IsSuccess: true,
	}
}

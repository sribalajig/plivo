package validator

import (
	"plivo/plivo.sms/model"
)

type TextValidator struct {
}

func (textValidator TextValidator) Validate(sms model.SMS) ValidationResult {
	if len(sms.Text) == 0 {
		return ValidationResult{
			Error: "text is missing",
		}
	}

	if len(sms.Text) > 120 || len(sms.Text) < 1 {
		return ValidationResult{
			Error: "text is invalid",
		}
	}

	return ValidationResult{
		IsSuccess: true,
	}
}

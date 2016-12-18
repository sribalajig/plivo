package validator

import (
	"plivo/plivo.sms/model"
)

type FromNumberValidator struct {
}

func (fromNumberValidator FromNumberValidator) Validate(sms model.SMS) ValidationResult {
	if len(sms.From) == 0 {
		return ValidationResult{
			Message: "from is missing",
		}
	}

	if !validate(sms.From) {
		return ValidationResult{
			Message: "from number is invalid",
		}
	}

	return ValidationResult{
		IsSuccess: true,
	}
}

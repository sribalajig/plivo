package validator

import (
	"plivo/plivo.sms/model"
)

type SMSValidator struct {
	validators []Validator
}

func (smsValidator SMSValidator) Validate(sms model.SMS) ValidationResult {
	for _, validator := range smsValidator.validators {
		result := validator.Validate(sms)

		if !result.IsSuccess {
			return result
		}
	}

	return ValidationResult{
		IsSuccess: true,
	}
}

func NewSMSValidator() SMSValidator {
	return SMSValidator{
		validators: []Validator{FromNumberValidator{}, ToNumberValidator{}},
	}
}

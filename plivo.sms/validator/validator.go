package validator

import (
	"plivo/plivo.sms/model"
)

type Validator interface {
	Validate(model.SMS) ValidationResult
}

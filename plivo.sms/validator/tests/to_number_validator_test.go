package tests

import (
	"plivo/plivo.sms/model"
	"plivo/plivo.sms/validator"
	"testing"
)

var toNumberTests = []struct {
	ToNumber string
	IsValid  bool
}{
	{"", false},
	{"123", false},
	{"123456", true},
	{"1234567891234567", true},
	{"12345678912345678", false},
}

func TestToNumberInSMS(t *testing.T) {
	toNumberValidator := validator.ToNumberValidator{}

	for _, testCase := range toNumberTests {
		validationResult := toNumberValidator.Validate(model.NewSMS("", testCase.ToNumber, ""))

		if testCase.IsValid != validationResult.IsSuccess {
			t.Errorf("To Number : %s, expected validation result is %t, actual is %t", testCase.ToNumber, testCase.IsValid, validationResult.IsSuccess)
			t.Fail()
		}
	}
}

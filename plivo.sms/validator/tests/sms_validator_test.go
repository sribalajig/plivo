package tests

import (
	"plivo/plivo.sms/model"
	"plivo/plivo.sms/validator"
	"testing"
)

var testCases = []struct {
	FromNumber string
	ToNumber   string
	Text       string
	IsValid    bool
}{
	{"", "", "", false},
	{"123", "123", "Hello", false},
	{"123456", "123456", "", false},
	{"1234567891234567", "123456", "", false},
	{"12345678912345678", "", "", false},
	{"123456", "123456", "hello world", true},
}

func TestSMSValidator(t *testing.T) {
	smsValidator := validator.NewSMSValidator()

	for _, testCase := range testCases {
		sms := model.NewSMS(testCase.FromNumber, testCase.ToNumber, testCase.Text)
		validationResult := smsValidator.Validate(sms)

		if testCase.IsValid != validationResult.IsSuccess {
			t.Errorf("SMS: %q, expected validation result is %t, actual is %t", sms, testCase.IsValid, validationResult.IsSuccess)
			t.Fail()
		}
	}
}

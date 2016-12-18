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
	{"123456", "123456", "", true},
	{"1234567891234567", "123456", "", false},
	{"12345678912345678", "", "", false},
}

func TestSMSValidator(t *testing.T) {
	fromNumberValidator := validator.FromNumberValidator{}

	for _, testCase := range testCases {
		validationResult := fromNumberValidator.Validate(model.NewSMS(testCase.FromNumber, testCase.ToNumber, testCase.Text))

		if testCase.IsValid != validationResult.IsSuccess {
			t.Errorf("From Number : %s, expected validation result is %t, actual is %t", testCase.FromNumber, testCase.IsValid, validationResult.IsSuccess)
			t.Fail()
		}
	}
}

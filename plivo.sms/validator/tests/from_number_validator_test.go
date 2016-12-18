package tests

import (
	"plivo/plivo.sms/model"
	"plivo/plivo.sms/validator"
	"testing"
)

var fromNumberTests = []struct {
	FromNumber string
	IsValid    bool
}{
	{"", false},
	{"123", false},
	{"123456", true},
	{"1234567891234567", true},
	{"12345678912345678", false},
}

func TestFromNumberInSMS(t *testing.T) {
	fromNumberValidator := validator.FromNumberValidator{}

	for _, testCase := range fromNumberTests {
		validationResult := fromNumberValidator.Validate(model.NewSMS(testCase.FromNumber, "", ""))

		if testCase.IsValid != validationResult.IsSuccess {
			t.Errorf("FromNumber : %s, expected validation result was %t, actual is %t", testCase.FromNumber, testCase.IsValid, validationResult.IsSuccess)
			t.Fail()
		}
	}
}

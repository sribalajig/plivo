package tests

import (
	"plivo/plivo.sms/model"
	"plivo/plivo.sms/validator"
	"testing"
)

var textTests = []struct {
	Text    string
	IsValid bool
}{
	{"", false},
	{"H", true},
	{"George Bernard Dantzig (November 8, 1914 – May 13, 2005) was an American mathematical scientist who made important contributions to operations research, computer science, economics, and statistics.", false},
}

func TestTextInSMS(t *testing.T) {
	textValidator := validator.TextValidator{}

	for _, testCase := range textTests {
		validationResult := textValidator.Validate(model.NewSMS("", "", testCase.Text))

		if testCase.IsValid != validationResult.IsSuccess {
			t.Errorf("Text : %s, expected validation result is %t, actual is %t", testCase.Text, testCase.IsValid, validationResult.IsSuccess)
			t.Log(validationResult.Message)
			t.Fail()
		}
	}
}

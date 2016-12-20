package tests

import (
	"plivo/plivo.sms/model"
	"testing"
)

var smsTextTests = []struct {
	Text    string
	IsValid bool
}{
	{"STOP", true},
	{"STOP\n", true},
	{"STOP\r\n", true},
	{"STOPIT", false},
	{"abcde STOP", false},
	{"wefwfdf", false},
	{"stop", false},
}

func TestSMSText(t *testing.T) {
	for _, testCase := range smsTextTests {
		sms := model.NewSMS("", "", testCase.Text)

		if stop := sms.Contains("STOP"); stop != testCase.IsValid {
			t.Errorf("For %s, expected %t but got %t", testCase.Text, testCase.IsValid, stop)
			t.Fail()
		}
	}
}

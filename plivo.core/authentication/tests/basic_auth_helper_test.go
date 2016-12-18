package tests

import (
	"plivo/plivo.core/authentication"
	"testing"
)

var testCases = []struct {
	header         string
	expectedResult *authentication.Credentials
	isSuccess      bool
}{
	{"", nil, false},
	{"Basic aGVsbG86d29ybGQ=", &authentication.Credentials{UserName: "hello", Password: "world"}, true},
}

func TestBasicAuthHelper(t *testing.T) {
	for _, testCase := range testCases {
		isSuccess, result := authentication.GetCredentials(testCase.header)

		if isSuccess != testCase.isSuccess {
			t.Errorf("For %s, expected %t, but got %t", testCase.header, testCase.isSuccess, isSuccess)
			t.Fail()
		}

		if testCase.expectedResult != nil {
			if result == nil {
				t.Error("Expected not nil credentials but got nil")
				t.Fail()
			} else if result.UserName != testCase.expectedResult.UserName {
				t.Errorf("Expected %s as user name but got %s", testCase.expectedResult.UserName, result.UserName)
				t.Fail()
			} else if result.Password != testCase.expectedResult.Password {
				t.Errorf("Expected %s as password but got %s", testCase.expectedResult.Password, result.Password)
				t.Fail()
			}
		}
	}
}

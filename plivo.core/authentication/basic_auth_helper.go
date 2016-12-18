package authentication

import (
	"encoding/base64"
	"strings"
)

func GetCredentials(header string) (bool, *Credentials) {
	auth := strings.SplitN(header, " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		return false, nil
	}

	creds, err := base64.StdEncoding.DecodeString(auth[1])

	if err != nil {
		return false, nil
	}

	pair := strings.SplitN(string(creds), ":", 2)

	if len(pair) != 2 {
		return false, nil
	}

	return true, &Credentials{
		UserName: pair[0],
		Password: pair[1],
	}
}

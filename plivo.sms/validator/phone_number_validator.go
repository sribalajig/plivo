package validator

func validate(phoneNumber string) bool {
	if len(phoneNumber) < 6 {
		return false
	}

	if len(phoneNumber) > 16 {
		return false
	}

	return true
}

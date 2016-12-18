package model

type InboundSMS struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

func (inboundSMS InboundSMS) Validate() bool {
	if !inboundSMS.validatePhoneNumber(inboundSMS.From) {
		return false
	}

	if !inboundSMS.validatePhoneNumber(inboundSMS.To) {
		return false
	}

	if len(inboundSMS.Text) == 0 {
		return false
	}

	return true
}

func (inboundSMS InboundSMS) validatePhoneNumber(phoneNumber string) bool {
	if len(phoneNumber) == 0 {
		return false
	}

	if len(phoneNumber) <= 6 {
		return false
	}

	if len(phoneNumber) >= 16 {
		return false
	}

	return true
}

var NewInboundSMS = func(from string, to string, text string) InboundSMS {
	return InboundSMS{
		From: from,
		To:   to,
		Text: text,
	}
}

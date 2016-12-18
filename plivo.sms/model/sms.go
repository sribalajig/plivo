package model

type SMS struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

func (sms SMS) Contains(pattern string) bool {
	return true
}

var NewSMS = func(from string, to string, text string) SMS {
	return SMS{
		From: from,
		To:   to,
		Text: text,
	}
}

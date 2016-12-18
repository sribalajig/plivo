package model

type SMS struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

var NewSMS = func(from string, to string, text string) SMS {
	return SMS{
		From: from,
		To:   to,
		Text: text,
	}
}

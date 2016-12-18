package repository

import (
	"plivo/plivo.core/phone_numbers/model"
)

type PhoneNumberRepository struct {
}

func (phoneNumberRepository PhoneNumberRepository) Get(number string) (model.PhoneNumber, error) {
	return model.PhoneNumber{}, nil
}

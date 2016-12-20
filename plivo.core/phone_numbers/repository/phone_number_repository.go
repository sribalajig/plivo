package repository

import (
	"plivo/plivo.persistence/postgre"
)

type PhoneNumberRepository struct {
}

func (phoneNumberRepository PhoneNumberRepository) Exists(number string) bool {
	return postgre.Exists(number)
}

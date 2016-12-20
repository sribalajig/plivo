package model

import (
	"plivo/plivo.core/accounts/model"
)

type PhoneNumber struct {
	Id      int
	Number  string         `orm:"column(number)"`
	Account *model.Account `orm:"rel(fk)"`
}

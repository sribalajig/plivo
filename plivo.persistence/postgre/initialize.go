package postgre

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"plivo/plivo.core/accounts/model"
	pn "plivo/plivo.core/phone_numbers/model"
)

func Initialize() error {
	orm.RegisterDriver("postgres", orm.DRPostgres)

	err := orm.RegisterDataBase("default", "postgres", "user=postgres password=01061988 host=localhost port=5432 sslmode=disable dbname=plivo")

	if err != nil {
		return err
	}

	account := new(model.Account)

	orm.RegisterModel(account)

	phoneNumber := new(pn.PhoneNumber)

	orm.RegisterModel(phoneNumber)

	return nil
}

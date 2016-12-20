package postgre

import (
	"fmt"
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

	acc := new(model.Account)

	orm.RegisterModel(acc)

	phoneNumber := new(pn.PhoneNumber)

	orm.RegisterModel(phoneNumber)

	o := orm.NewOrm()

	qs := o.QueryTable(acc)

	var account model.Account

	fmt.Println("this is the name")

	qs.Filter("username", "plivo1").One(&account)

	fmt.Println(account)

	return nil
}

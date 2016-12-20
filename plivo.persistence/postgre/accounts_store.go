package postgre

import (
	"github.com/astaxie/beego/orm"
	"plivo/plivo.core/accounts/model"
)

func GetAccount(userName string) model.Account {
	o := orm.NewOrm()

	qs := o.QueryTable(new(model.Account))

	var account model.Account

	qs.Filter("username", "plivo1").One(&account)

	return account
}

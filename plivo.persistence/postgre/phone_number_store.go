package postgre

import (
	"github.com/astaxie/beego/orm"
	"plivo/plivo.core/phone_numbers/model"
)

func Exists(number string) bool {
	return orm.NewOrm().QueryTable(new(model.PhoneNumber)).Filter("number", number).Exist()
}

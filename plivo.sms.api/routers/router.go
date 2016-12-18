// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"plivo/plivo.sms.api/controllers"

	"github.com/astaxie/beego"

	"plivo/plivo.sms.api/filters"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/inbound/sms", &controllers.SMSController{}, "post:InBouund")

	beego.InsertFilter("/inbound/sms", beego.BeforeRouter, filters.FilterInboundSMS)
}
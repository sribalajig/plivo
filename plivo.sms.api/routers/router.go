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
	beego.Router("/inbound/sms", &controllers.SMSController{}, "post:Inbound")
	beego.Router("/outbound/sms", &controllers.SMSController{}, "post:Outbound")

	beego.InsertFilter("/*", beego.BeforeRouter, filters.BasicAuthFilter)

	beego.InsertFilter("/*/sms", beego.BeforeRouter, filters.FilterSMS)

	beego.InsertFilter("/inbound/sms", beego.BeforeRouter, filters.FilterInboundSMS)
	beego.InsertFilter("/outbound/sms", beego.BeforeRouter, filters.FilterOutboundSMS)
}

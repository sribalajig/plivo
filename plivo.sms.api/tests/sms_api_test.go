package test

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"plivo/plivo.persistence/postgre"
	"plivo/plivo.persistence/redis"
	_ "plivo/plivo.sms.api/routers"
	"runtime"
	"testing"
)

func init() {
	if postgre.Initialize() != nil {
		fmt.Println("Failed to inititalize Postgre. Exiting...")

		return
	}

	redis.Initialize()

	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestInboundSMS(t *testing.T) {
	var jsonStr = []byte(`{
	  "to" : "4924195509192",
	  "from" : "61871112940",
	  "text" : "Hello"
	}`)

	r, _ := http.NewRequest("POST", "/inbound/sms", bytes.NewBuffer(jsonStr))

	r.Header.Add("Authorization", "Basic cGxpdm8xOjIwUzBLUE5PSU0=")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestInboundSMS", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Inbound SMS Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// TestGet is a sample to run an endpoint test
func TestOutboundSMS(t *testing.T) {
	var jsonStr = []byte(`{
	  "to" : "4924195509192",
	  "from" : "61871112940",
	  "text" : "Hello"
	}`)

	r, _ := http.NewRequest("POST", "/outbound/sms", bytes.NewBuffer(jsonStr))

	r.Header.Add("Authorization", "Basic cGxpdm8xOjIwUzBLUE5PSU0=")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestOutboundSMS", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Outbound SMS Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

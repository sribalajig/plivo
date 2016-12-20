package tests

import (
	"plivo/plivo.persistence/redis"
	"plivo/plivo.sms/model"
	"plivo/plivo.sms/repository"
	"testing"
)

func TestSetStop(t *testing.T) {
	redis.Initialize()

	repository.SMSRepository{}.SetStop(model.NewSMS("123456", "234567", "STOP"))

	val, _ := redis.Get("from:234567;to:123456")

	if val != "stop" {
		t.Errorf("Exptected stop, found %s", val)

		t.Fail()
	}

	redis.Delete("from:234567;to:123456")
}

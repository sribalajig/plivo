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

func TestUpdateSentSMSCount(t *testing.T) {
	redis.Initialize()

	repository.SMSRepository{}.UpdateSentSMSCount(model.NewSMS("123456", "234567", "Hello"))
	repository.SMSRepository{}.UpdateSentSMSCount(model.NewSMS("123456", "456789", "hi"))
	repository.SMSRepository{}.UpdateSentSMSCount(model.NewSMS("123456", "987654", "hi there"))

	val, _ := redis.Get("from:123456")

	if val != "3" {
		t.Errorf("Expected %s, got %s", "2", val)
	}

	redis.Delete("from:123456")
}

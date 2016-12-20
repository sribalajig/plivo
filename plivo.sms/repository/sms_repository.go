package repository

import (
	"fmt"
	"plivo/plivo.persistence/redis"
	"plivo/plivo.sms/model"
	"strconv"
	"time"
)

type SMSRepository struct {
}

func (smsRepository SMSRepository) SetStop(sms model.SMS) error {
	key := fmt.Sprintf("from:%s;to:%s", sms.To, sms.From)

	return redis.Set(key, "stop", time.Hour*4)
}

func (smsRepository SMSRepository) IsStopped(sms model.SMS) (bool, error) {
	key := fmt.Sprintf("from:%s;to:%s", sms.From, sms.To)

	result, _ := redis.Get(key)

	if result == "stop" {
		return true, nil
	}

	return false, nil
}

func (smsRepository SMSRepository) GetSentSMSCount(sms model.SMS) (int, error) {
	key := fmt.Sprintf("from:%s", sms.From)

	val, err := redis.Get(key)

	if val == "" {
		return 0, err
	}

	return strconv.Atoi(val)
}

func (smsRepository SMSRepository) UpdateSentSMSCount(sms model.SMS) error {
	key := fmt.Sprintf("from:%s", sms.From)

	count, err := redis.Increment(key)

	if count == 1 {
		redis.Expire(key, time.Hour*24)
	}

	return err
}

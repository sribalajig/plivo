package repository

import (
	"plivo/plivo.sms/model"
)

type SMSRepository struct {
}

func (smsRepository SMSRepository) SetStop(sms model.SMS) error {
	return nil
}

func (smsRepository SMSRepository) IsStopped(sms model.SMS) (bool, error) {
	return true, nil
}

func (smsRepository SMSRepository) GetSentSMSCount(sms model.SMS) (int, error) {
	return 1, nil
}

func (smsRepository SMSRepository) UpdateSentSMSCount(sms model.SMS) error {
	return nil
}

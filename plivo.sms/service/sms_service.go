package service

import (
	"plivo/plivo.sms/model"
	"plivo/plivo.sms/repository"
)

type SMSService struct {
	smsRepository repository.SMSRepository
}

func NewSMSService() SMSService {
	return SMSService{
		smsRepository: repository.SMSRepository{},
	}
}

func (smsService SMSService) ProcessInboundSMS(sms model.SMS) error {
	if sms.Contains("STOP") {
		return smsService.smsRepository.SetStop(sms)
	}

	return nil
}

func (smsService SMSService) ProcessOutboundSMS(sms model.SMS) error {
	return smsService.smsRepository.UpdateSentSMSCount(sms)
}

func (smsService SMSService) IsStopped(sms model.SMS) (bool, error) {
	return smsService.smsRepository.IsStopped(sms)
}

func (smsService SMSService) GetSentSMSCount(sms model.SMS) (int, error) {
	return smsService.smsRepository.GetSentSMSCount(sms)
}

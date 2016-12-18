package service

import (
	"fmt"
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
	if isStopped, _ := smsService.smsRepository.IsStopped(sms); isStopped {
		return fmt.Errorf("sms from %s to %s blocked by STOP request", sms.From, sms.To)
	}

	if count, _ := smsService.smsRepository.GetSentSMSCount(sms); count > 50 {
		return fmt.Errorf("limit reached for from %s", sms.From)
	}

	return smsService.smsRepository.UpdateSentSMSCount(sms)
}

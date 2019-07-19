package services

type SMSService interface {
	Send()
}

type TwilioSMSService struct {
}

func (s *TwilioSMSService) Send() {

}

func (s *TwilioSMSService) Webhook() {

}

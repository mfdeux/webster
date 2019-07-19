package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v3"
)

type EmailService interface {
	Send()
}

// MailGunEmailClient is an email client
type MailGunEmailClient struct {
	client *mailgun.MailgunImpl
}

// NewMailGunClient creates a new mailgun client
func NewMailGunClient(domain, apiKey string) *MailGunEmailClient {
	mg := mailgun.NewMailgun(domain, apiKey)
	return &MailGunEmailClient{
		client: mg,
	}
}

// Send sends an email message
func (mg *MailGunEmailClient) Send() {
	sender := "sender@example.com"
	subject := "Fancy subject!"
	body := "Hello from Mailgun Go!"
	recipient := "recipient@example.com"

	// The message object allows you to add attachments and Bcc recipients
	message := mg.client.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message	with a 10 second timeout
	resp, id, err := mg.client.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}

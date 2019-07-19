package services

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

// StripePaymentService is a service
type StripePaymentService struct {
	client *client.API
}

// NewStripePaymentService is
func NewStripePaymentService() *StripePaymentService {
	sc := &client.API{}
	sc.Init("sk_key", nil)
	return &StripePaymentService{
		client: sc,
	}
}

// SavePaymentMethod saves a payment method
func (s *StripePaymentService) SavePaymentMethod() (string, error) {
	args := &stripe.ChargeParams{
		Amount:      stripe.Int64(2000),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String("Charge for jenny.rosen@example.com"),
	}
	args.SetSource("tok_amex") // obtained with Stripe.js
	args.SetIdempotencyKey("QubES0gV8mcZEHpY")
	charge, err := s.client.Charges.New(args)
	if err != nil {
		return "", nil
	}
	return charge.ID, err
}

// Charge charges a new payment method
func (s *StripePaymentService) Charge() (string, error) {
	args := &stripe.ChargeParams{
		Amount:      stripe.Int64(2000),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String("Charge for jenny.rosen@example.com"),
	}
	args.SetSource("tok_amex") // obtained with Stripe.js
	args.SetIdempotencyKey("QubES0gV8mcZEHpY")
	charge, err := s.client.Charges.New(args)
	if err != nil {
		return "", nil
	}
	return charge.ID, err
}

// Refund provides a refund
func (s *StripePaymentService) Refund() (string, error) {
	args := &stripe.ChargeParams{
		Amount:      stripe.Int64(2000),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String("Charge for jenny.rosen@example.com"),
	}
	args.SetSource("tok_amex") // obtained with Stripe.js
	args.SetIdempotencyKey("QubES0gV8mcZEHpY")
	charge, err := s.client.Charges.New(args)
	if err != nil {
		return "", nil
	}
	return charge.ID, err
}

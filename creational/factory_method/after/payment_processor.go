package main

import (
	"fmt"
)

type PaymentProcessor interface {
	Pay(amount int) string
}

type PaypalPaymentProcessor struct {
	APIKey string
}

func (p *PaypalPaymentProcessor) Pay(amount int) string {
	return fmt.Sprintf("Processed $%d payment through PayPal", amount)
}

type StripePaymentProcessor struct {
	SecretKey string
	Region    string
}

func (s *StripePaymentProcessor) Pay(amount int) string {
	return fmt.Sprintf("Processed $%d payment through Stripe", amount)
}

type PaymentFactory interface {
	CreatePaymentProcessor() PaymentProcessor
}

type PaypalPaymentFactory struct{}

func (f *PaypalPaymentFactory) CreatePaymentProcessor() PaymentProcessor {
	return &PaypalPaymentProcessor{APIKey: "PAYPAL_KEY"}
}

type StripePaymentFactory struct{}

func (f *StripePaymentFactory) CreatePaymentProcessor() PaymentProcessor {
	return &StripePaymentProcessor{SecretKey: "STRIPE_KEY", Region: "us_east_1"}
}

func CreatePaymentFactory(paymentType string) PaymentFactory {
	switch paymentType {
	case "paypal":
		return &PaypalPaymentFactory{}
	case "stripe":
		return &StripePaymentFactory{}
	default:
		return nil
	}
}

func ProcessPayment(factory PaymentFactory, amount int) string {
	processor := factory.CreatePaymentProcessor()
	return processor.Pay(amount)
}

func main() {
	var factory = CreatePaymentFactory("paypal")
	fmt.Println(ProcessPayment(factory, 100))

	factory = CreatePaymentFactory("stripe")
	fmt.Println(ProcessPayment(factory, 200))
}

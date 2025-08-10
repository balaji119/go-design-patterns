package main

import (
	"fmt"
)

type PaymentProcessor interface {
	ProcessPayment(amount int) string
}

type PaypalPaymentProcessor struct {
	APIKey string
}

func (p *PaypalPaymentProcessor) ProcessPayment(amount int) string {
	return fmt.Sprintf("Processed $%d payment through PayPal", amount)
}

type StripePaymentProcessor struct {
	SecretKey string
	Region    string
}

func (s *StripePaymentProcessor) ProcessPayment(amount int) string {
	return fmt.Sprintf("Processed $%d payment through Stripe", amount)
}

func CreatePaymentProcessor(paymentType string, amount int) (PaymentProcessor, error) {
	var processor PaymentProcessor
	switch paymentType {
	case "paypal":
		processor = &PaypalPaymentProcessor{APIKey: "PAYPAL_KEY"}
	case "stripe":
		processor = &StripePaymentProcessor{SecretKey: "STRIPES_KEY", Region: "us_east_1"}
	default:
		return nil, fmt.Errorf("Invalid payment type")
	}
	return processor, nil
}

func Process(paymentType string, amount int) {
	processor, err := CreatePaymentProcessor(paymentType, amount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(processor.ProcessPayment(amount))
}

func main() {
	Process("paypal", 100)
	Process("stripe", 200)
}

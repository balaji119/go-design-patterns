package main

type EmailMessage struct {
	To                 string
	Cc                 string
	Bcc                string
	Subject            string
	Body               string
	Attachments        []string
	RequestReadReceipt bool
	Priority           string
}

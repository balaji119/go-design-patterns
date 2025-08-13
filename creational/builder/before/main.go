package main

import (
	"fmt"
)

func main() {

	email := EmailMessage{
		To:                 "recipient@example.com",
		Cc:                 "cc@example.com",
		Bcc:                "bcc@example.com",
		Subject:            "Hello",
		Body:               "This is a test email.",
		Attachments:        []string{"file1.txt", "file2.txt"},
		RequestReadReceipt: true,
		Priority:           "High",
	}

	fmt.Println(email)
}

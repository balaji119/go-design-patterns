package main

import (
	"fmt"
)

func main() {

	builder := EmailBuilder{}

	email := builder.SetTo("recipient@example.com").
		SetCc("cc@example.com").
		SetBcc("bcc@example.com").
		SetSubject("Hello").
		SetBody("This is a test email.").
		AddAttachment("file.txt").
		SetRequestReadReceipt(true).
		SetPriority("High").
		Build()

	fmt.Println(email)
}

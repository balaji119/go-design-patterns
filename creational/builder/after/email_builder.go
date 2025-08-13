package main

type EmailBuilder struct {
	email EmailMessage
}

func (e *EmailBuilder) SetTo(to string) *EmailBuilder {
	e.email.To = to
	return e
}

func (e *EmailBuilder) SetCc(cc string) *EmailBuilder {
	e.email.Cc = cc
	return e
}

func (e *EmailBuilder) SetBcc(bcc string) *EmailBuilder {
	e.email.Bcc = bcc
	return e
}

func (e *EmailBuilder) SetSubject(subject string) *EmailBuilder {
	e.email.Subject = subject
	return e
}

func (e *EmailBuilder) SetBody(body string) *EmailBuilder {
	e.email.Body = body
	return e
}

func (e *EmailBuilder) AddAttachment(attachment string) *EmailBuilder {
	e.email.Attachments = append(e.email.Attachments, attachment)
	return e
}

func (e *EmailBuilder) SetRequestReadReceipt(request bool) *EmailBuilder {
	e.email.RequestReadReceipt = request
	return e
}

func (e *EmailBuilder) SetPriority(priority string) *EmailBuilder {
	e.email.Priority = priority
	return e
}

func (e *EmailBuilder) Build() EmailMessage {
	if e.email.To == "" {
		panic("Email 'To' field is required")
	}
	if e.email.Subject == "" {
		panic("Email 'Subject' field is required")
	}
	return e.email
}

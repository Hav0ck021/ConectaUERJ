package model

type GmailSender struct {
	Name              string
	FromEmailAddress  string
	FromEmailPassword string
}

type EmailService interface {
	SendEmail(
		subject string,
		content string,
		to []string,
	) error
}

package service

import (
	"github.com/OVillas/user-api/config"
	"github.com/OVillas/user-api/model"
	"net/smtp"
	"strconv"
)

type emailService struct {
	gmailSender model.GmailSender
}

func NewEmailService(name string, fromEmailAddress string, fromEmailPassword string) model.EmailService {
	gmailSender := model.GmailSender{
		Name:              name,
		FromEmailAddress:  fromEmailAddress,
		FromEmailPassword: fromEmailPassword,
	}

	return &emailService{
		gmailSender: gmailSender,
	}
}

func (sender *emailService) SendEmail(subject string, content string, to []string) error {
	message := []byte("Subject: " + subject + "\r\n" +
		"From: " + sender.gmailSender.Name + " <" + sender.gmailSender.FromEmailAddress + ">\r\n" +
		"To: " + to[0] + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		content)

	auth := smtp.PlainAuth("", sender.gmailSender.FromEmailAddress, sender.gmailSender.FromEmailPassword, config.SMTPServer)

	smtpPortStr := strconv.Itoa(config.SMTPPort)

	err := smtp.SendMail(config.SMTPServer+":"+smtpPortStr, auth, sender.gmailSender.FromEmailAddress, to, message)
	if err != nil {
		return err
	}

	return nil
}

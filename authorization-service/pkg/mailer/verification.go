package mailer

import (
	"fmt"
	"strconv"

	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/config"

	"gopkg.in/gomail.v2"
)

type Interface interface {
	DialAndSend(userEmail string, userId int32,
		verificationToken string) error
}

type Verification struct {
	cfg *config.Config
}

func NewVerification(cfg *config.Config) *Verification {
	return &Verification{
		cfg: cfg,
	}
}

func (v Verification) DialAndSend(userEmail string, userId int32,
	verificationToken string) error {

	m := gomail.NewMessage()
	id := strconv.Itoa(int(userId))

	m.SetHeader("From", v.cfg.Mailer.AuthEmail)
	m.SetHeader("To", userEmail)

	url := v.cfg.Mailer.CallBackHost + "?" + "user-id=" + id + "&" + "token=" +
		verificationToken
	message := "Please click the following link to verify your email:\r\n" +
		url

	m.SetHeader("Subject", "Email verification")
	m.SetBody("text/html", message)

	d := gomail.NewDialer(v.cfg.Mailer.SmtpHost, v.cfg.Mailer.Port,
		v.cfg.Mailer.AuthName, v.cfg.Mailer.AuthPassword)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("verification - DialAndSend - DialAndSend: %w", err)
	}

	return nil
}

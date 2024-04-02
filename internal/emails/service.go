package emails

import (
	_ "embed"
	"fmt"

	"github.com/MateoCaicedoW/GO-SMTP/email"
	"github.com/MateoCaicedoW/GO-SMTP/server"
	"github.com/MateoCaicedoW/email-sender/internal/app/config"
)

type SenderService interface {
	SendEmail(name, message, to string, attachments []email.Attachment) error
}

var _ SenderService = (*service)(nil)

type service struct {
	client *server.SMTPServer
}

func NewService(client *server.SMTPServer) *service {
	return &service{client: client}
}

//go:embed email.html
var codeVerification string

func (s *service) SendEmail(name, message, to string, attachments []email.Attachment) error {
	email := email.Params{
		SenderName:      name,
		Sender:          "smtp.gmail.com",
		To:              []string{to},
		Subject:         "Email Sender",
		Body:            fmt.Sprintf(codeVerification, fmt.Sprintf("%s/show", config.BaseURL), message),
		BodyContentType: "text/html",
		Attachments:     attachments,
	}

	err := email.Send(s.client)
	if err != nil {
		return err
	}

	return nil
}

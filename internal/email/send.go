package email

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/MateoCaicedoW/GO-SMTP/email"
	"github.com/MateoCaicedoW/GO-SMTP/server"
)

//go:embed email.html
var codeVerification string

func Send(s *server.SMTPServer, name, message, to string, attachments []email.Attachment) error {
	email := email.Params{
		SenderName:      name,
		Sender:          "smtp.gmail.com",
		To:              []string{to},
		Subject:         "Email Sender",
		Body:            fmt.Sprintf(codeVerification, fmt.Sprintf("%s/show", os.Getenv("BASE_URL")), message),
		BodyContentType: "text/html",
		Attachments:     attachments,
	}

	err := email.Send(s)
	if err != nil {
		return err
	}

	return nil
}

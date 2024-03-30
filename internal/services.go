package internal

import (
	myServer "github.com/MateoCaicedoW/GO-SMTP/server"
	"github.com/MateoCaicedoW/email-sender/internal/app/config"
	"github.com/MateoCaicedoW/email-sender/internal/email"
	"github.com/leapkit/core/server"
)

func AddServices(r server.Router) error {
	s := myServer.NewSMTP("smtp.gmail.com", "587", config.SenderEmail, config.SenderPassword, "")
	r.Use(server.InCtxMiddleware("mailerService", email.NewService(s)))

	return nil
}

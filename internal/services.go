package internal

import (
	myServer "github.com/MateoCaicedoW/GO-SMTP/server"
	"github.com/MateoCaicedoW/email-sender/internal/app/config"
	"github.com/MateoCaicedoW/email-sender/internal/companies"
	"github.com/MateoCaicedoW/email-sender/internal/emails"
	"github.com/MateoCaicedoW/email-sender/internal/sender"
	"github.com/MateoCaicedoW/email-sender/internal/subscribers"
	"github.com/MateoCaicedoW/email-sender/internal/users"
	"github.com/leapkit/core/server"
)

func AddServices(r server.Router) error {
	s := myServer.NewSMTP("smtp.gmail.com", "587", config.SenderEmail, config.SenderPassword, "")
	conn, err := DB()
	if err != nil {
		return err
	}

	r.Use(server.InCtxMiddleware("mailerService", sender.NewService(s)))
	r.Use(server.InCtxMiddleware("subscriberService", subscribers.NewService(conn)))
	r.Use(server.InCtxMiddleware("emailService", emails.NewService(conn)))
	r.Use(server.InCtxMiddleware("userService", users.NewService(conn)))
	r.Use(server.InCtxMiddleware("companyService", companies.NewService(conn)))

	return nil
}

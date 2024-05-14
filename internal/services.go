package internal

import (
	myServer "github.com/MateoCaicedoW/GO-SMTP/server"

	"github.com/MateoCaicedoW/email-sender/internal/sender"
	"github.com/MateoCaicedoW/email-sender/internal/users"
	"github.com/leapkit/core/server"
)

func AddServices(r server.Router) error {
	s := myServer.NewSMTP("smtp.gmail.com", "587", "smtpmateo@gmail.com", "rdej kqnl pczk ixve", "")
	conn, err := DB()
	if err != nil {
		return err
	}

	r.Use(server.InCtxMiddleware("mailerService", sender.NewService(s)))
	r.Use(server.InCtxMiddleware("userService", users.NewService(conn)))

	return nil
}

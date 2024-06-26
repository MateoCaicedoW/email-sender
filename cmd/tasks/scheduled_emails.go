package tasks

import (
	"fmt"
	"time"

	"github.com/MateoCaicedoW/GO-SMTP/email"
	myServer "github.com/MateoCaicedoW/GO-SMTP/server"
	"github.com/MateoCaicedoW/email-sender/internal"
	"github.com/MateoCaicedoW/email-sender/internal/app/config"
	"github.com/MateoCaicedoW/email-sender/internal/emails"
	"github.com/MateoCaicedoW/email-sender/internal/sender"
	"github.com/MateoCaicedoW/email-sender/internal/subscribers"
)

func ScheduledEmails() error {
	conn, err := internal.DB()
	if err != nil {
		fmt.Println(err)
		return err
	}
	s := myServer.NewSMTP("smtp.gmail.com", "587", config.SenderEmail, config.SenderPassword, "")

	service := emails.NewService(conn)
	subService := subscribers.NewService(conn)
	mailer := sender.NewService(s)
	emails, err := service.Scheduled()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, e := range emails {
		if e.ScheduledAt.After(time.Now()) {
			continue
		}

		fmt.Printf("Sending email... %s\n %s\n", e.Name, e.Subject)
		var attachments email.Attachments
		if e.AttachmentName != "" {
			attachments = append(attachments, email.Attachment{
				FileName: e.AttachmentName,
				Content:  e.AttachmentContent,
			})
		}

		subs, err := subService.All(e.CompanyID)
		if err != nil {
			fmt.Println(err)
			return err
		}

		for _, sub := range subs {

			err := mailer.SendEmail(e.Name, e.Message, sub.Email, e.Subject, attachments)
			if err != nil {
				fmt.Println(err)
				return err
			}

			e.Sent = true
			if err := service.Update(&e); err != nil {
				fmt.Println(err)
				return err
			}

		}

		fmt.Println("Email Sent ✅")
	}
	return nil
}

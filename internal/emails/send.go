package emails

import (
	"io"
	"net/http"

	"github.com/MateoCaicedoW/GO-SMTP/email"
	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/MateoCaicedoW/email-sender/internal/sender"
	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/form"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/session"
)

func Send(w http.ResponseWriter, r *http.Request) {
	s := r.Context().Value("mailerService").(sender.SenderService)
	subService := r.Context().Value("subscriberService").(models.SubscriberService)
	emailService := r.Context().Value("emailService").(models.EmailService)
	session := session.FromCtx(r.Context())
	companyID := session.Values["company_id"].(uuid.UUID)
	em := models.Email{}
	if err := form.Decode(r, &em); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	em.CompanyID = companyID
	verrs := emailService.Validate(&em)
	if verrs.HasAny() {
		rx := render.FromCtx(r.Context())

		rx.Set("errors", verrs.Errors)
		rx.Set("email", em)
		if err := rx.RenderWithLayout("emails/new.html", "modal.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	file, header, err := r.FormFile("attachment")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var attachments []email.Attachment

	if file != nil {
		defer file.Close()

		bytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		attachments = append(attachments, email.Attachment{
			FileName: header.Filename,
			Content:  bytes,
		})

		em.AttachmentName = header.Filename
		em.AttachmentContent = attachments[0].Content
	}

	if !em.Scheduled {
		subs, err := subService.All(companyID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, sub := range subs {
			if err := s.SendEmail(em.Name, em.Message, sub.Email, em.Subject, attachments); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		em.Sent = true
	}

	if err := emailService.Create(&em); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/emails/")
}

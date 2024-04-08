package emails

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/session"
)

func Indicators(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	emailService := r.Context().Value("emailService").(models.EmailService)
	session := session.FromCtx(r.Context())
	companyID := session.Values["company_id"].(uuid.UUID)
	emailSent, err := emailService.CountSent(companyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	emailScheduled, err := emailService.CountScheduled(companyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rx.Set("emailsSent", emailSent)
	rx.Set("emailsScheduled", emailScheduled)
	if err := rx.RenderClean("emails/indicators.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

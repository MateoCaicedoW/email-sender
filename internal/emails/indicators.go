package emails

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/leapkit/core/render"
)

func Indicators(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	emailService := r.Context().Value("emailService").(models.EmailService)

	emailSent, err := emailService.CountSent()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	emailScheduled, err := emailService.CountScheduled()
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

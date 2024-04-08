package emails

import (
	"net/http"
	"time"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/leapkit/core/render"
)

func New(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	email := models.Email{
		ScheduledAt: time.Now(),
	}

	rx.Set("email", email)
	if err := rx.RenderWithLayout("emails/new.html", "modal.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

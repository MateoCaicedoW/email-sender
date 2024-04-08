package subscribers

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/session"
)

func Indicators(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	subService := r.Context().Value("subscriberService").(models.SubscriberService)
	session := session.FromCtx(r.Context())
	companyID := session.Values["company_id"].(uuid.UUID)
	subscribersCount, err := subService.ActiveCount(companyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	unsubscribersCount, err := subService.InactiveCount(companyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rx.Set("unsubscribersCount", unsubscribersCount)
	rx.Set("subscribersCount", subscribersCount)
	if err := rx.RenderClean("subscribers/indicators.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

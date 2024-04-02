package subscribers

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/leapkit/core/render"
)

func Indicators(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	subService := r.Context().Value("subscriberService").(models.SubscriberService)
	subscribersCount, err := subService.ActiveCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	unsubscribersCount, err := subService.InactiveCount()
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

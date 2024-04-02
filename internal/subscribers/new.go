package subscribers

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"

	"github.com/leapkit/core/form"
	"github.com/leapkit/core/render"
)

func New(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	sub := &models.Subscriber{}

	rx.Set("subscriber", sub)
	if err := rx.RenderWithLayout("subscribers/new.html", "modal.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	subService := r.Context().Value("subscriberService").(models.SubscriberService)
	sub := &models.Subscriber{}

	if err := form.Decode(r, sub); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	errs := subService.Validate(sub)
	if errs.HasAny() {
		rx := render.FromCtx(r.Context())

		rx.Set("subscriber", sub)
		rx.Set("errors", errs.Errors)
		if err := rx.RenderWithLayout("subscribers/new.html", "modal.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	if err := subService.Create(sub); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/subscribers")
}

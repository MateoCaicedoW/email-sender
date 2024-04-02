package subscribers

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/form"
	"github.com/leapkit/core/render"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	subService := r.Context().Value("subscriberService").(models.SubscriberService)
	subID := uuid.FromStringOrNil(r.PathValue("id"))
	sub, err := subService.Find(subID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rx.Set("subscriber", sub)
	if err := rx.RenderWithLayout("subscribers/edit.html", "modal.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	subService := r.Context().Value("subscriberService").(models.SubscriberService)
	subID := uuid.FromStringOrNil(r.PathValue("id"))
	sub, err := subService.Find(subID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := form.Decode(r, sub); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sub.ID = subID
	errs := subService.Validate(sub)
	if errs.HasAny() {
		rx := render.FromCtx(r.Context())

		rx.Set("subscriber", sub)
		rx.Set("errors", errs.Errors)
		if err := rx.RenderWithLayout("subscribers/edit.html", "modal.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	if err := subService.Update(sub); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/subscribers")
}

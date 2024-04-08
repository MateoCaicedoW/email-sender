package subscribers

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gofrs/uuid/v5"

	"github.com/leapkit/core/form"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/session"
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
	session := session.FromCtx(r.Context())
	companyID := session.Values["company_id"].(uuid.UUID)

	if err := form.Decode(r, sub); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sub.CompanyID = companyID
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

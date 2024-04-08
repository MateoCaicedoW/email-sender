package subscribers

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/session"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	subService := r.Context().Value("subscriberService").(models.SubscriberService)
	subID := uuid.FromStringOrNil(r.PathValue("id"))
	session := session.FromCtx(r.Context())
	companyID := session.Values["company_id"].(uuid.UUID)
	sub, err := subService.Find(subID, companyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := subService.Delete(sub.ID, companyID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/subscribers")
}

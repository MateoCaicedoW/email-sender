package users

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gofrs/uuid/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	userService := r.Context().Value("userService").(models.UserService)

	id := uuid.FromStringOrNil(r.PathValue("id"))
	// session := session.FromCtx(r.Context())
	// currentUser := session.Values["user_id"].(uuid.UUID)
	// companyID := session.Values["company_id"].(uuid.UUID)
	err := userService.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/")

}

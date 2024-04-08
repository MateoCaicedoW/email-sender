package users

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/session"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	userService := r.Context().Value("userService").(models.UserService)
	companyService := r.Context().Value("companyService").(models.CompanyService)
	subID := uuid.FromStringOrNil(r.PathValue("id"))
	session := session.FromCtx(r.Context())
	currentUser := session.Values["user_id"].(uuid.UUID)
	companyID := session.Values["company_id"].(uuid.UUID)
	user, err := userService.FindByID(subID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := companyService.RemoveUser(companyID, user.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if currentUser != user.ID {
		w.Header().Set("HX-Redirect", "/users")
		return
	}

	firstCompany, err := userService.FirstCompany(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if firstCompany.ID.IsNil() {
		session.Values["company_id"] = nil
		session.Save(r, w)
		w.Header().Set("HX-Redirect", "/companies/new")
		return
	}

	session.Values["company_id"] = firstCompany.ID
	session.Save(r, w)

	w.Header().Set("HX-Redirect", "/emails")

}

package users

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
	user := &models.User{}

	rx.Set("user", user)
	if err := rx.RenderWithLayout("users/new.html", "modal.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	userService := r.Context().Value("userService").(models.UserService)
	user := models.User{}
	session := session.FromCtx(r.Context())
	companyID := session.Values["company_id"].(uuid.UUID)

	if err := form.Decode(r, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	errs := userService.Validate(user, companyID)
	if errs.HasAny() {
		rx := render.FromCtx(r.Context())

		rx.Set("user", user)
		rx.Set("errors", errs.Errors)
		if err := rx.RenderWithLayout("users/new.html", "modal.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	u, err := userService.FindByEmail(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	companyService := r.Context().Value("companyService").(models.CompanyService)
	if !u.ID.IsNil() {
		if err := companyService.AddUser(companyID, u.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Redirect", "/users")
		return
	}

	if err := userService.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := companyService.AddUser(companyID, user.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/users")
}

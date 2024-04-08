package auth

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/session"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	userService := r.Context().Value("userService").(models.UserService)
	email := r.FormValue("Email")

	if email == "" {
		rx.Set("email", email)
		rx.Set("errors", map[string]string{"email": "Email is required."})
		if err := rx.RenderWithLayout("auth/login.html", "auth.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	user, err := userService.FindByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.ID.IsNil() {
		rx.Set("email", email)
		rx.Set("errors", map[string]string{"email": "User not found, please register first or try again."})
		if err := rx.RenderWithLayout("auth/login.html", "auth.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return

	}
	session := session.FromCtx(r.Context())
	session.Values["user_id"] = user.ID
	session.Save(r, w)

	company, err := userService.FirstCompany(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if company.ID.IsNil() {
		http.Redirect(w, r, "/companies/new", http.StatusFound)
		return
	}

	session.Values["company_id"] = company.ID
	session.Save(r, w)

	http.Redirect(w, r, "/emails", http.StatusFound)
}

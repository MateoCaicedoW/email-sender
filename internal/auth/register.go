package auth

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/leapkit/core/form"
	"github.com/leapkit/core/render"
)

func Register(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	user := models.User{}
	rx.Set("user", user)
	if err := rx.RenderWithLayout("auth/register.html", "auth.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	userService := r.Context().Value("userService").(models.UserService)
	rx := render.FromCtx(r.Context())
	if err := form.Decode(r, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	verrs := userService.ValidateRegister(user)
	if verrs.HasAny() {
		rx.Set("user", user)
		rx.Set("errors", verrs.Errors)
		if err := rx.RenderWithLayout("auth/register.html", "auth.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := userService.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

package users

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/models"
	"github.com/leapkit/core/form"
	"github.com/leapkit/core/render"
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

	if err := form.Decode(r, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	errs := userService.Validate(user)
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

	if err := userService.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/")
}

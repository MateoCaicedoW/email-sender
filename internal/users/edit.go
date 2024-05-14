package users

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/form"
	"github.com/leapkit/core/render"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	userService := r.Context().Value("userService").(models.UserService)
	userID := uuid.FromStringOrNil(r.PathValue("id"))

	user, err := userService.FindByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rx.Set("user", user)
	if err := rx.RenderWithLayout("users/edit.html", "modal.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	userService := r.Context().Value("userService").(models.UserService)
	userID := uuid.FromStringOrNil(r.PathValue("id"))
	user, err := userService.FindByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := form.Decode(r, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	errs := userService.Validate(user)
	if errs.HasAny() {
		rx := render.FromCtx(r.Context())

		rx.Set("user", user)
		rx.Set("errors", errs.Errors)
		if err := rx.RenderWithLayout("users/edit.html", "modal.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	if err := userService.Update(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/")
}

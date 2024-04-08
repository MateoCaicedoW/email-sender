package auth

import (
	"net/http"

	"github.com/leapkit/core/render"
)

func Login(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	rx.Set("email", "")
	if err := rx.RenderWithLayout("auth/login.html", "auth.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

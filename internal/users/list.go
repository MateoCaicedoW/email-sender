package users

import (
	"net/http"
	"strconv"

	"github.com/MateoCaicedoW/email-sender/internal/models"
	"github.com/leapkit/core/render"
)

func List(w http.ResponseWriter, r *http.Request) {
	userService := r.Context().Value("userService").(models.UserService)
	rx := render.FromCtx(r.Context())

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	term := r.URL.Query().Get("term")
	status := r.URL.Query().Get("status")

	usrs, err := userService.List(10, page, term)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rx.Set("list", usrs)
	rx.Set("term", term)
	rx.Set("page", page)
	rx.Set("status", status)
	if r.Header.Get("HX-Request") == "true" {

		w.Header().Add("HX-Push-Url", r.URL.String())

		if err := rx.RenderClean("users/table.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	if err := rx.Render("users/list.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

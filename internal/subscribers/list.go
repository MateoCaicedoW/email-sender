package subscribers

import (
	"net/http"
	"strconv"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/leapkit/core/render"
)

func List(w http.ResponseWriter, r *http.Request) {
	subService := r.Context().Value("subscriberService").(models.SubscriberService)
	rx := render.FromCtx(r.Context())

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	term := r.URL.Query().Get("term")
	status := r.URL.Query().Get("status")

	subs, err := subService.List(10, page, term, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rx.Set("list", subs)
	rx.Set("term", term)
	rx.Set("page", page)
	rx.Set("status", status)
	if r.Header.Get("HX-Request") == "true" {

		w.Header().Add("HX-Push-Url", r.URL.String())

		if err := rx.RenderClean("subscribers/table.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	if err := rx.Render("subscribers/list.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

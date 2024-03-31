package subscribers

import (
	"net/http"

	"github.com/leapkit/core/render"
)

func List(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())

	if err := rx.Render("subscribers/list.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

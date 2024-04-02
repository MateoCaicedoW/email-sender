package emails

import (
	"net/http"

	"github.com/leapkit/core/render"
)

func List(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())

	if err := rx.Render("emails/list.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

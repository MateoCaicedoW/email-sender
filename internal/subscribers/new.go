package subscribers

import (
	"net/http"

	"github.com/leapkit/core/render"
)

func New(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())

	if err := rx.RenderWithLayout("subscribers/new.html", "modal.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

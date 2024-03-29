package actions

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderWithLayout(w, "/home.html", "application.html")
}

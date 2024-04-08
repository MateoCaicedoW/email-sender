package home

import (
	"net/http"

	"github.com/leapkit/core/session"
)

func Index(w http.ResponseWriter, r *http.Request) {
	session := session.FromCtx(r.Context())
	userID := session.Values["user_id"]

	if userID == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/emails", http.StatusSeeOther)
}

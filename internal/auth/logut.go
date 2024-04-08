package auth

import (
	"net/http"

	"github.com/leapkit/core/session"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session := session.FromCtx(r.Context())
	session.Values["user_id"] = nil
	session.Values["company_id"] = nil
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

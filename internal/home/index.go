package home

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/emails", http.StatusSeeOther)
}

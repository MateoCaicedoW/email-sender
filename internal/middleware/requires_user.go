package middleware

import (
	"context"
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/models"
	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/session"
)

func RequiresUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := session.FromCtx(r.Context())

		userID := session.Values["user_id"]
		if userID == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		id := userID.(uuid.UUID)

		user := models.User{
			ID:        id,
			FirstName: "John",
			LastName:  "Doe",
		}

		rx := render.FromCtx(r.Context())
		rx.Set("currentUser", user)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "currentUser", user)))
	})
}

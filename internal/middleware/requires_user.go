package middleware

import (
	"context"
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/leapkit/core/render"
)

func RequiresUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := models.User{
			FirstName: "John",
			LastName:  "Smith",
		}

		rx := render.FromCtx(r.Context())
		rx.Set("currentUser", user)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "currentUser", user)))
	})
}

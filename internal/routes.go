package internal

import (
	"cmp"
	"os"

	"github.com/MateoCaicedoW/email-sender/internal/app/helpers"
	"github.com/MateoCaicedoW/email-sender/internal/users"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/server"
	"github.com/leapkit/core/session"
)

// AddRoutes mounts the routes for the application,
// it assumes that the base services have been injected
// in the creation of the server instance.
func AddRoutes(r server.Router) error {
	// LeapKit Middleware
	r.Use(session.Middleware(
		cmp.Or(os.Getenv("SESSION_SECRET"), "d720c059-9664-4980-8169-1158e167ae57"),
		cmp.Or(os.Getenv("SESSION_NAME"), "leapkit_session"),
	))

	r.Use(render.Middleware(
		render.TemplateFS(tmpls, "internal"),

		render.WithDefaultLayout("layout.html"),
		render.WithHelpers(helpers.All),
	))

	r.Group("/", func(rx server.Router) {
		rx.HandleFunc("GET /{$}", users.List)
		rx.HandleFunc("GET /new", users.New)
		rx.HandleFunc("POST /{$}", users.Create)
		rx.HandleFunc("GET /{id}/edit", users.Edit)
		rx.HandleFunc("PUT /{id}", users.Update)
		rx.HandleFunc("DELETE /{id}", users.Delete)
		rx.HandleFunc("POST /{id}/send_email", users.SendEmail)
	})

	// Mounting the assets manager at the end of the routes
	// so that it can serve the public assets.
	r.HandleFunc(Assets.HandlerPattern(), Assets.HandlerFn)

	return nil
}

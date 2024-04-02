package internal

import (
	"github.com/MateoCaicedoW/email-sender/internal/app/helpers"
	"github.com/MateoCaicedoW/email-sender/internal/email"
	"github.com/MateoCaicedoW/email-sender/internal/home"
	"github.com/MateoCaicedoW/email-sender/internal/subscribers"
	"github.com/leapkit/core/envor"
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
		envor.Get("SESSION_SECRET", "d720c059-9664-4980-8169-1158e167ae57"),
		envor.Get("SESSION_NAME", "leapkit_session"),
	))

	r.Use(render.Middleware(
		render.TemplateFS(tmpls, "internal"),

		render.WithDefaultLayout("layout.html"),
		render.WithHelpers(helpers.All),
	))

	r.HandleFunc("GET /emails", home.Index)
	r.HandleFunc("GET /{$}", home.Index)
	r.HandleFunc("POST /send_email", email.Send)
	r.HandleFunc("GET /show_email", email.Show)

	r.Group("/subscribers/", func(r server.Router) {
		r.HandleFunc("GET /{$}", subscribers.List)
		r.HandleFunc("GET /new", subscribers.New)
		r.HandleFunc("POST /{$}", subscribers.Create)
		r.HandleFunc("GET /{id}/edit", subscribers.Edit)
		r.HandleFunc("PUT /{id}", subscribers.Update)
		r.HandleFunc("DELETE /{id}", subscribers.Delete)
		r.HandleFunc("GET /indicators", subscribers.Indicators)
	})

	// Mounting the assets manager at the end of the routes
	// so that it can serve the public assets.
	r.HandleFunc(Assets.HandlerPattern(), Assets.HandlerFn)

	return nil
}

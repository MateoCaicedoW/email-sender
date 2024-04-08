package internal

import (
	"github.com/MateoCaicedoW/email-sender/internal/app/helpers"
	"github.com/MateoCaicedoW/email-sender/internal/auth"
	"github.com/MateoCaicedoW/email-sender/internal/companies"
	"github.com/MateoCaicedoW/email-sender/internal/emails"
	"github.com/MateoCaicedoW/email-sender/internal/home"
	"github.com/MateoCaicedoW/email-sender/internal/middleware"
	"github.com/MateoCaicedoW/email-sender/internal/subscribers"
	"github.com/MateoCaicedoW/email-sender/internal/users"
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
	r.HandleFunc("GET /{$}", home.Index)
	// auth routes
	r.HandleFunc("GET /login", auth.Login)
	r.HandleFunc("POST /login", auth.Authenticate)
	r.HandleFunc("GET /register/{$}", auth.Register)
	r.HandleFunc("POST /register", auth.SignUp)

	r.Group("/companies/", func(r server.Router) {
		r.HandleFunc("GET /new", companies.New)
		r.HandleFunc("POST /new", companies.Create)
	})

	// private routes
	r.Group("/", func(r server.Router) {
		r.Use(middleware.RequiresUser)
		r.Use(middleware.RequiresCompany)

		r.HandleFunc("GET /switch_company/{company_id}", companies.Switch)

		r.Group("/emails/", func(rx server.Router) {
			rx.HandleFunc("GET /{$}", emails.List)
			rx.HandleFunc("GET /new", emails.New)
			rx.HandleFunc("POST /send", emails.Send)
			rx.HandleFunc("GET /indicators", emails.Indicators)
		})

		r.Group("/subscribers/", func(rx server.Router) {
			rx.HandleFunc("GET /{$}", subscribers.List)
			rx.HandleFunc("GET /new", subscribers.New)
			rx.HandleFunc("POST /{$}", subscribers.Create)
			rx.HandleFunc("GET /{id}/edit", subscribers.Edit)
			rx.HandleFunc("PUT /{id}", subscribers.Update)
			rx.HandleFunc("DELETE /{id}", subscribers.Delete)
			rx.HandleFunc("GET /indicators", subscribers.Indicators)
		})

		r.Group("/users/", func(rx server.Router) {
			rx.HandleFunc("GET /{$}", users.List)
			rx.HandleFunc("GET /new", users.New)
			rx.HandleFunc("POST /{$}", users.Create)
			rx.HandleFunc("GET /{id}/edit", users.Edit)
			rx.HandleFunc("PUT /{id}", users.Update)
			rx.HandleFunc("DELETE /{id}", users.Delete)
		})

		r.HandleFunc("GET /logout", auth.Logout)
	})

	// Mounting the assets manager at the end of the routes
	// so that it can serve the public assets.
	r.HandleFunc(Assets.HandlerPattern(), Assets.HandlerFn)

	return nil
}

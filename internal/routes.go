package internal

import (
	"cmp"
	"os"

	"github.com/MateoCaicedoW/email-sender/internal/middleware"
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

	r.Use(middleware.RequiresUser)

	// Routes

	// Mounting the assets manager at the end of the routes
	// so that it can serve the public assets.
	r.HandleFunc(Assets.HandlerPattern(), Assets.HandlerFn)

	return nil
}

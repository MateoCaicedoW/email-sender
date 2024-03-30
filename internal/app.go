package internal

import (
	"embed"

	"github.com/MateoCaicedoW/email-sender/public"
	"github.com/leapkit/core/assets"
	"github.com/leapkit/core/db"
	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/gloves"

	"github.com/paganotoni/tailo"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed **/*.html **/*.html *.html
	tmpls embed.FS

	// Assets is the manager for the public assets
	// it allows to watch for changes and reload the assets
	// when changes are made.
	Assets = assets.NewManager(public.Files)

	// TailoOptions allow to define how to compile
	// the tailwind css files, which is the input and
	// what will be the output.
	TailoOptions = []tailo.Option{
		tailo.UseInputPath("internal/assets/application.css"),
		tailo.UseOutputPath("public/application.css"),
		tailo.UseConfigPath("tailwind.config.js"),
	}

	// GlovesOptions are the options that will be used by the gloves
	// tool to hot reload the application.
	GlovesOptions = []gloves.Option{
		// Run the tailo watcher so when changes are made to
		// the html code it rebuilds css.
		gloves.WithRunner(tailo.WatcherFn(TailoOptions...)),
		gloves.WithRunner(Assets.Watch),
		gloves.WatchExtension(".go", ".css", ".js"),
	}

	// DatabaseURL to connect and interact with our database instance.
	DatabaseURL = envor.Get("DATABASE_URL", "leapkit.db")
	// DB is the database connection builder function
	// that will be used by the application based on the driver and
	// connection string.
	DB = db.ConnectionFn(DatabaseURL, db.WithDriver("sqlite3"))
)

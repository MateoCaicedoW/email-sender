package internal

import (
	"embed"
	"time"

	"github.com/MateoCaicedoW/email-sender/public"

	"github.com/leapkit/core/assets"

	"github.com/leapkit/core/form"
	"github.com/leapkit/core/gloves"

	"github.com/paganotoni/tailo"

	_ "github.com/lib/pq"
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
		gloves.WatchExtension(".go", ".css", ".js", ".html"),
	}
)

func init() {
	//register custom type decoder functions for time.Time
	form.RegisterCustomTypeFunc(DecodeTime, time.Time{})
}

func DecodeTime(vals []string) (interface{}, error) {
	if len(vals) == 0 {
		return time.Time{}, nil
	}

	val := vals[0]
	if val == "" {
		return time.Time{}, nil
	}

	tt, err := time.Parse("2006-01-02T15:04:05", val)
	if err != nil {
		return nil, err
	}

	return tt, nil
}

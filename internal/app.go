package internal

import (
	"embed"
	"time"

	"github.com/MateoCaicedoW/email-sender/public"

	"github.com/leapkit/core/assets"

	"github.com/leapkit/core/form"
	"github.com/leapkit/core/gloves"

	"github.com/paganotoni/tailo"
)

var (
	//go:embed **/*.html **/*.html *.html
	tmpls embed.FS

	Assets = assets.NewManager(public.Files)

	TailoOptions = []tailo.Option{
		tailo.UseInputPath("internal/assets/application.css"),
		tailo.UseOutputPath("public/application.css"),
		tailo.UseConfigPath("tailwind.config.js"),
	}

	GlovesOptions = []gloves.Option{
		gloves.WithRunner(tailo.WatcherFn(TailoOptions...)),
		gloves.WithRunner(Assets.Watch),
		gloves.WatchExtension("", ".css", ".js", ""),
	}
)

func init() {
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

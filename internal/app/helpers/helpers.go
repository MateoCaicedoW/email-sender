package helpers

import (
	"regexp"

	"github.com/leapkit/core/hctx"
	"github.com/leapkit/core/render"
)

var All = hctx.Merge(
	render.AllHelpers,
	hctx.Map{
		"activeClass": activeClass,
	},
)

func activeClass(pattern, currentUrl string) string {
	if matched, err := regexp.MatchString(pattern, currentUrl); !matched || err != nil {
		return ""
	}

	return "active"
}

package main

import (
	"fmt"

	"github.com/MateoCaicedoW/email-sender/internal"
	"github.com/leapkit/core/gloves"
)

func main() {
	err := gloves.Start(
		"cmd/app/main.go",

		internal.GlovesOptions...,
	)

	if err != nil {
		fmt.Println(err)
	}
}

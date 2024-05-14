package main

import (
	"cmp"
	"fmt"
	"net/http"
	"os"

	"github.com/MateoCaicedoW/email-sender/internal"
	_ "github.com/leapkit/core/envload"
	"github.com/leapkit/core/server"
)

func main() {
	s := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3000")),
	)

	if err := internal.AddServices(s); err != nil {
		os.Exit(1)
	}

	if err := internal.AddRoutes(s); err != nil {
		os.Exit(1)
	}

	fmt.Println("Server started at", s.Addr())
	err := http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println(err)
	}
}

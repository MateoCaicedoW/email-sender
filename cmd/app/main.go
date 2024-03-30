package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MateoCaicedoW/email-sender/internal"
	"github.com/leapkit/core/envor"
	"github.com/leapkit/core/server"
)

func main() {
	s := server.New(
		server.WithHost(envor.Get("HOST", "127.0.0.1")),
		server.WithPort(envor.Get("PORT", "3000")),
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

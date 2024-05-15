package main

import (
	"fmt"

	"github.com/MateoCaicedoW/email-sender/internal"
	"github.com/MateoCaicedoW/email-sender/internal/migrations"
	"github.com/leapkit/core/db"
	"github.com/paganotoni/tailo"
)

func main() {
	// Setup tailo to compile tailwind css.
	err := tailo.Setup()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("✅ Tailwind CSS setup successfully")

	if err := db.Create(internal.DatabaseURL); err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("✅ Database created successfully")

	conn, err := internal.DB()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.RunMigrations(migrations.All, conn)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("✅ Migrations ran successfully")
}

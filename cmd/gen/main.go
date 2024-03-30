package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	usage := func() {
		fmt.Println("Usage: gen migration <name>")
		fmt.Println("Available generators:")
		fmt.Println("  -  migration")
	}

	if len(os.Args) < 3 {
		usage()

		return
	}

	switch os.Args[1] {
	case "migration":

		migrationName := fmt.Sprintf("%s_%s.sql", os.Args[2], time.Now().Format("20060102150405"))
		fmt.Println("Generating migration", migrationName)

		filepath := filepath.Join("db/migrations", migrationName)

		file, err := os.Create(filepath)
		if err != nil {
			fmt.Println("Error creating migration file:", err)
			return
		}

		defer file.Close()

		file.WriteString("-- +migrate Up\n\n\n-- +migrate Down\n\n")

		fmt.Println("Migration generated successfully")
		return
	default:
		usage()
	}

}

package internal

import (
	"cmp"
	"os"

	"github.com/leapkit/core/db"
	_ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

var (
	DatabaseURL = cmp.Or(os.Getenv("DATABASE_URL"), "postgres://postgres:postgres@localhost:5432/email_sender?sslmode=disable")

	DB = db.ConnectionFn(DatabaseURL)

	// DB = db.ConnectionFn(DatabaseURL, db.WithDriver("sqlite3"))
)

package internal

import (
	"cmp"
	"os"

	"github.com/leapkit/core/db"
)

var (
	DatabaseURL = cmp.Or(os.Getenv("DATABASE_URL"), "postgres://postgres:postgres@localhost:5432/email_sender?sslmode=disable")

	DB = db.ConnectionFn(DatabaseURL)
)

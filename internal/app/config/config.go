package config

import "github.com/leapkit/core/envor"

var (
	BaseURL        = envor.Get("BASE_URL", "")
	SenderEmail    = envor.Get("SENDER_EMAIL", "")
	SenderPassword = envor.Get("SENDER_PASSWORD", "")
)

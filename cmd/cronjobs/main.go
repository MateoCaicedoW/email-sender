package main

import (
	"fmt"
	"log"
	"time"

	"github.com/MateoCaicedoW/email-sender/cmd/tasks"
	"github.com/go-co-op/gocron"
)

func main() {
	loc, err := time.LoadLocation("America/Bogota")
	if err != nil {
		log.Fatal(fmt.Errorf("error loading time zone: %w", err))
	}

	engine := gocron.NewScheduler(loc)

	jobs := []struct {
		schedule string
		task     func() error
	}{
		{"*/2 * * * *", tasks.ScheduledEmails},
	}

	for _, job := range jobs {
		_, err := engine.Cron(job.schedule).Do(job.task)
		if err != nil {
			log.Fatal(fmt.Errorf("error scheduling job: %w", err))
		}
	}

	log.Println("[Info] Starting cron worker")
	engine.StartBlocking()
}

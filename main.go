package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MateoCaicedoW/GO-SMTP/server"
	"github.com/MateoCaicedoW/email-sender/actions"
	"github.com/MateoCaicedoW/email-sender/render"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	server := mux.NewRouter()
	server.Use(setServices)
	server.HandleFunc("/", actions.Home).Methods("GET")
	server.HandleFunc("/send", actions.SendEmail).Methods("POST")
	server.HandleFunc("/show", actions.ShowEmail).Methods("GET")

	srv := &http.Server{
		Handler:      server,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server running on port 3000")
	log.Fatal(srv.ListenAndServe())
}

func setServices(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := server.NewSMTP("smtp.gmail.com", "587", os.Getenv("SENDER_EMAIL"), os.Getenv("SENDER_PASSWORD"), "")

		r = r.WithContext(context.WithValue(r.Context(), "smtp", s))
		render.SetData("smtp", s)
		next.ServeHTTP(w, r)
	})
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

package users

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/MateoCaicedoW/email-sender/internal/sender"

	"github.com/gofrs/uuid/v5"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	s := r.Context().Value("mailerService").(sender.SenderService)
	userService := r.Context().Value("userService").(models.UserService)

	id := uuid.FromStringOrNil(r.PathValue("id"))
	user, err := userService.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.SendEmail("New Email", "Random Message", user.Email, "Test Email", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/")
}

package emails

import (
	"net/http"

	"github.com/MateoCaicedoW/GO-SMTP/email"
	"github.com/MateoCaicedoW/email-sender/internal/app/models"
)

func Send(w http.ResponseWriter, r *http.Request) {
	s := r.Context().Value("mailerService").(*service)
	subService := r.Context().Value("subscriberService").(models.SubscriberService)
	message := r.FormValue("message")
	// file, header, err := r.FormFile("attachment")
	// if err != nil && err != http.ErrMissingFile {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	var attachments []email.Attachment

	// if file != nil {
	// 	defer file.Close()

	// 	bytes, err := io.ReadAll(file)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}

	// 	attachments = append(attachments, email.Attachment{
	// 		FileName: header.Filename,
	// 		Content:  bytes,
	// 	})
	// }

	subs, err := subService.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, sub := range subs {
		if err := s.SendEmail("Sender App", message, sub.Email, attachments); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("HX-Redirect", "/emails/")
}

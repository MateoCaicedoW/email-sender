package email

import (
	"fmt"
	"io"
	"net/http"

	"github.com/MateoCaicedoW/GO-SMTP/email"
)

func Send(w http.ResponseWriter, r *http.Request) {
	s := r.Context().Value("mailerService").(*service)
	e := r.FormValue("email")
	name := r.FormValue("name")
	message := r.FormValue("message")
	file, header, err := r.FormFile("attachment")
	if err != nil && err != http.ErrMissingFile {
		fmt.Println("Error getting file 1")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var attachments []email.Attachment

	if file != nil {
		defer file.Close()

		bytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading file")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		attachments = append(attachments, email.Attachment{
			FileName: header.Filename,
			Content:  bytes,
		})
	}

	if err := s.SendEmail(name, message, e, attachments); err != nil {
		fmt.Println("Error sending email")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

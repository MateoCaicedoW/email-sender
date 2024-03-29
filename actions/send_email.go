package actions

import (
	"fmt"
	"io"
	"net/http"

	"github.com/MateoCaicedoW/GO-SMTP/email"
	"github.com/MateoCaicedoW/GO-SMTP/server"
	myEmail "github.com/MateoCaicedoW/email-sender/internal/email"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	s := r.Context().Value("smtp").(*server.SMTPServer)
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

	err = myEmail.Send(s, name, message, e, attachments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

package emails

import (
	"net/http"
	"strconv"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/session"
)

func List(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	emailService := r.Context().Value("emailService").(models.EmailService)
	session := session.FromCtx(r.Context())
	companyID := session.Values["company_id"].(uuid.UUID)
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	term := r.URL.Query().Get("term")
	status := r.URL.Query().Get("status")

	emails, err := emailService.List(10, page, term, status, companyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rx.Set("list", emails)
	rx.Set("term", term)
	rx.Set("page", page)
	rx.Set("status", status)

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Add("HX-Push-Url", r.URL.String())

		if err := rx.RenderClean("emails/table.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	if err := rx.Render("emails/list.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

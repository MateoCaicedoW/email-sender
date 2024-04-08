package middleware

import (
	"context"
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/session"
)

func RequiresCompany(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := session.FromCtx(r.Context())

		companyID := session.Values["company_id"]
		if companyID == nil {
			http.Redirect(w, r, "/companies/new", http.StatusSeeOther)
			return
		}

		userID := session.Values["user_id"]
		if userID == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		id := companyID.(uuid.UUID)
		companyService := r.Context().Value("companyService").(models.CompanyService)

		currentCompany, err := companyService.FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rx := render.FromCtx(r.Context())
		rx.Set("currentCompany", currentCompany)

		companies, err := companyService.AllByUser(userID.(uuid.UUID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rx.Set("userCompanies", companies)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "currentCompany", currentCompany)))
	})
}

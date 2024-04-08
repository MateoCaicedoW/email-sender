package companies

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gofrs/uuid/v5"
	"github.com/leapkit/core/session"
)

func Switch(w http.ResponseWriter, r *http.Request) {
	companyID := uuid.FromStringOrNil(r.PathValue("company_id"))
	companyService := r.Context().Value("companyService").(models.CompanyService)
	company, err := companyService.FindByID(companyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if company.ID.IsNil() {
		http.Redirect(w, r, "/emails", http.StatusFound)
		return
	}

	session := session.FromCtx(r.Context())
	session.Values["company_id"] = company.ID
	session.Save(r, w)

	http.Redirect(w, r, "/emails", http.StatusFound)
}

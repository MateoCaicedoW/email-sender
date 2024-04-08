package companies

import (
	"net/http"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gofrs/uuid/v5"

	"github.com/leapkit/core/form"
	"github.com/leapkit/core/render"
	"github.com/leapkit/core/session"
)

func New(w http.ResponseWriter, r *http.Request) {
	rx := render.FromCtx(r.Context())
	company := models.Company{}
	rx.Set("company", company)
	if err := rx.RenderWithLayout("companies/new.html", "auth.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	company := models.Company{}
	session := session.FromCtx(r.Context())
	if session.Values["user_id"] == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	companyService := r.Context().Value("companyService").(models.CompanyService)
	rx := render.FromCtx(r.Context())

	if err := form.Decode(r, &company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	verrs := companyService.Validate(company)
	if verrs.HasAny() {
		rx.Set("company", company)
		rx.Set("errors", verrs.Errors)
		if err := rx.RenderWithLayout("companies/new.html", "auth.html"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := companyService.Create(&company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := session.Values["user_id"].(uuid.UUID)
	err := companyService.AddUser(company.ID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["company_id"] = company.ID
	session.Save(r, w)

	http.Redirect(w, r, "/emails", http.StatusFound)
}

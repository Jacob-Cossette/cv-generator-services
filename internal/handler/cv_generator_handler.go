package handler

import (
	"html/template"
	"net/http"
)

type Handler struct {
}

type Proposition struct {
}

func (h *Handler) GenerateCV(w http.ResponseWriter, r *http.Request, propositions []Proposition) {
	// Parse your HTML template
	tmpl, err := template.ParseFiles("path/to/cv_template.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Proposals []Proposition
	}{
		Proposals: propositions,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

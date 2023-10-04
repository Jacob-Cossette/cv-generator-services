package handler

import (
	"fmt"
	"net/http"

	"github.com/jacob-cossette/cv-generator-services/internal/structure"
)

type SimilarityHandler struct {
	Job  structure.Job_offer
	User structure.User
}

func (h *SimilarityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The similarity between the user and the job is %f", h.calculateSimilarity())
}

func (h *SimilarityHandler) calculateSimilarity() float64 {
	similarity := 0.0
	for _, userSkill := range h.User.Skills {
		for _, jobSkill := range h.Job.Skills {
			if userSkill == jobSkill {
				similarity += 1.0
			}
		}
	}
	return similarity / float64(len(h.User.Skills))
}

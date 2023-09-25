type SimilarityHandler struct {
	Job  *Job
	User *User
}

func (h *SimilarityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The similarity between the user and the job is %f", h.calculateSimilarity())
}

func (h *SimilarityHandler) calculateSimilarity() float64 {
	similarity := 0.0
	for _, userSkill := range h.User.Skillset {
		for _, jobSkill := range h.Job.Skillset {
			if userSkill == jobSkill {
				similarity += 1.0
			}
		}
	}
	return similarity / float64(len(h.User.Skillset))
}
package structure

type Matcher_responses struct {
	UserID        string            `json:"user_id"`
	JobID         string            `json:"job_id"`
	MatchedSkills []string          `json:"matched_skills"`
	MatchedExp    int               `json:"matched_experience"`
	OtherInfo     map[string]string `json:"other_info"`
}

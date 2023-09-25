type job_offer struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	Skills      []string  `json:"skills"`
	Experience  string    `json:"experience"`
	Salary      float64   `json:"salary"`
	PostedAt    time.Time `json:"posted_at"`
}
type user struct {
	ID          string       `json:"id"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Email       string       `json:"email"`
	PhoneNumber string       `json:"phone_number"`
	Skills      []Skill      `json:"skills"`
	Experience  []Experience `json:"experience"`
}

type Experience struct {
	Title     string    `json:"title"`
	Company   string    `json:"company"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type Skill struct {
	Name        string `json:"name"`
	Proficiency string `json:"proficiency"`
	Description string `json:"description"`
	YearsOfExp  int    `json:"years_of_experience"`
}
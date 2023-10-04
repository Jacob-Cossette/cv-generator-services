package structure

type Skill struct {
	Name        string `json:"name"`
	Proficiency string `json:"proficiency"`
	Description string `json:"description"`
	YearsOfExp  int    `json:"years_of_experience"`
}

package structure

import "time"

type Profil struct {
	ID             string          `json:"id"`
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	Email          string          `json:"email"`
	PhoneNumber    string          `json:"phone_number"`
	Skills         []Skill         `json:"skills"`
	Experience     []Experience    `json:"experience"`
	Education      []Education     `json:"education"`
	Certifications []Certification `json:"certifications"`
	Languages      []Language      `json:"languages"`
	Interests      string          `json:"interests"`
	References     []Reference     `json:"references"`
	SocialMedia    SocialMedia     `json:"social_media"`
}

type SocialMedia struct {
	Facebook  string `json:"facebook,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	LinkedIn  string `json:"linkedin,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	GitHub    string `json:"github,omitempty"`
	Website   string `json:"website,omitempty"`
}
type Experience struct {
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}

type Education struct {
	Degree         string    `json:"degree"`
	School         string    `json:"school"`
	Location       string    `json:"location"`
	GraduationDate time.Time `json:"graduation_date"`
}

type Certification struct {
	Name         string    `json:"name"`
	Organization string    `json:"organization"`
	IssueDate    time.Time `json:"issue_date"`
}

type Language struct {
	Name        string `json:"name"`
	Proficiency string `json:"proficiency"`
}

type Reference struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

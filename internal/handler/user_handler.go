package handler

import (
	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacob-cossette/cv-generator-services/internal/structure"
)

func getSkills(c *gin.Context) {
	//c.IndentedJson(http.StatusOK, skills)
}

var Skills = []structure.Skill{
	{Name: "test Skill", Proficiency: "Test P", Description: "Test Description", YearsOfExp: 5},
	{Name: "test Skill", Proficiency: "Test P", Description: "Test Description", YearsOfExp: 5},
}

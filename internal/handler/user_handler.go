package handler

import (
	//"net/http"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacob-cossette/cv-generator-services/internal/structure"
)

func GetSkills(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Skills)
}

var Skills = []structure.Skill{
	{Name: "test Skill", Proficiency: "Test P", Description: "Test Description", YearsOfExp: 5},
	{Name: "test Skill", Proficiency: "Test P", Description: "Test Description", YearsOfExp: 5},
}

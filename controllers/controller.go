package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuellfa/study-go-gin/database"
	"github.com/samuellfa/study-go-gin/models"
)

func AllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, &students)
}

func FindStudentById(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, &student)
}

func FindStudentByCpf(c *gin.Context) {
	cpf := c.Params.ByName("cpf")

	var student models.Student

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, &student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateStudentPayload(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}
	database.DB.Delete(&student)
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateStudentPayload(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

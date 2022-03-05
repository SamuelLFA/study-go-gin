package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samuellfa/study-go-gin/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/students", controllers.AllStudents)
	r.GET("/students/:id", controllers.FindStudentById)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCpf)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.POST("/students", controllers.CreateStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.Run(":8080")
}

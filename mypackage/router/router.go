package router

import (
	"mypackage/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(studentController *controller.StudentController) *gin.Engine {

	r := gin.Default()

	r.POST("/students", studentController.CreateStudent)
	r.GET("/students", studentController.GetAllStudents)
	r.GET("/students/:id", studentController.GetStudentByID)
	r.PUT("/students/:id", studentController.UpdateStudent)
	r.DELETE("/students/:id", studentController.DeleteStudent)

	return r
}

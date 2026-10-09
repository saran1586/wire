package controller

import (
"net/http"
"strconv"

"mypackage/model"
"mypackage/service"

"github.com/gin-gonic/gin"
)

type StudentController struct {
Service *service.StudentService
}

func NewStudentController(service *service.StudentService) *StudentController {
return &StudentController{
Service: service,
}
}

func (c *StudentController) CreateStudent(ctx *gin.Context) {
var student model.Student

if err := ctx.ShouldBindJSON(&student); err != nil {
ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}

if err := c.Service.CreateStudent(&student); err != nil {
ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}

ctx.JSON(http.StatusCreated, student)
}

func (c *StudentController) GetAllStudents(ctx *gin.Context) {
students, err := c.Service.GetAllStudents()
if err != nil {
ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}

ctx.JSON(http.StatusOK, students)
}

func (c *StudentController) GetStudentByID(ctx *gin.Context) {
id, _ := strconv.Atoi(ctx.Param("id"))

student, err := c.Service.GetStudentByID(uint(id))
if err != nil {
ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
return
}

ctx.JSON(http.StatusOK, student)
}

func (c *StudentController) UpdateStudent(ctx *gin.Context) {
id, _ := strconv.Atoi(ctx.Param("id"))

var student model.Student

if err := ctx.ShouldBindJSON(&student); err != nil {
ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}

student.ID = uint(id)

if err := c.Service.UpdateStudent(&student); err != nil {
ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}
ctx.JSON(http.StatusOK, student)
}

func (c *StudentController) DeleteStudent(ctx *gin.Context) {
id, _ := strconv.Atoi(ctx.Param("id"))

if err := c.Service.DeleteStudent(uint(id)); err != nil {
ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}
ctx.JSON(http.StatusOK, gin.H{
"message": "Student deleted successfully",
})
}

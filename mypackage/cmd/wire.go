

package main

import (
	"mypackage/controller"
	"mypackage/database"
	"mypackage/repository"
	"mypackage/service"

	"github.com/google/wire"
)

func InitializeStudentController() (*controller.StudentController, error) {
	wire.Build(
database.NewDB,
repository.NewStudentRepository,
service.NewStudentService,
controller.NewStudentController,
)

return &controller.StudentController{}, nil
}
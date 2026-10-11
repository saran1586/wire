package main

import (
	"log"

	"mypackage/router"
)

func main() {

	studentController, err := InitializeStudentController()
	if err != nil {
		log.Fatal(err)
	}

	r := router.SetupRouter(studentController)

	log.Println("Server running on :8080")

	r.Run(":8080")
}

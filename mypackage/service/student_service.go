package service

import (
"mypackage/model"
"mypackage/repository"
)

type StudentService struct {
Repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
return &StudentService{
Repo: repo,
}
}

func (s *StudentService) CreateStudent(student *model.Student) error {
return s.Repo.CreateStudent(student)
}

func (s *StudentService) GetAllStudents() ([]model.Student, error) {
return s.Repo.GetAllStudents()
}

func (s *StudentService) GetStudentByID(id uint) (model.Student, error) {
return s.Repo.GetStudentByID(id)
}

func (s *StudentService) UpdateStudent(student *model.Student) error {
return s.Repo.UpdateStudent(student)
}

func (s *StudentService) DeleteStudent(id uint) error {
return s.Repo.DeleteStudent(id)
}

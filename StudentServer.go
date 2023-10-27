package main

import (
	"errors"
)

type StudentServer struct{}

func (s *StudentServer) GetStudentDetails(id int) (Student, error) {
	if id == 1 {
		return Student{"John Doe", 20, "Mathematics"}, nil
	}
	return Student{}, errors.New("Student not found")
}

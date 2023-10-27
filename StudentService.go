package main

type Student struct {
	Name    string
	Age     int
	Subject string
}

type StudentService interface {
	GetStudentDetails(id int) (Student, error)
}

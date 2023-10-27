package main

import "fmt"

func main() {
	server := &StudentServer{}

	// Simulating a client request to the ORB
	student, err := server.GetStudentDetails(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Student Details: %+v\n", student)
	}
}

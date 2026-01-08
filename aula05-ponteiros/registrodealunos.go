package main

import "fmt"

type Student struct {
	Name    string
	Surname string
	DNI     int
	Date    string
}

func (s *Student) detail() {
	fmt.Printf("ID: %d - %s %s - %d - %s\n", s.DNI, s.Name, s.Surname, s.DNI, s.Date)
}

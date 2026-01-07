package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (e *Employee) PrintEmployee() {
	fmt.Printf("ID: %d - %s - %s - %s\n", e.ID, e.Person.Name, e.Position, e.Person.DateOfBirth)
}

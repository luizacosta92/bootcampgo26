package main

import ( 
	"fmt"
)

func main() {
	fmt.Println("=== AULA 05 ===")

	student1 := Student{
		Name:    "João",
		Surname: "Silva",
		DNI:     1234567890,
		Date:    "2026-01-01",
	}
	student1.detail()

	student2 := &Student{"Maria", "Gomes", 1234567891, "2026-01-02"}
	student2.detail()

	var student3 Student
	student3.Name = "Pedro"
	student3.Surname = "Souza"
	student3.DNI = 1234567892
	student3.Date = "2026-01-03"
	student3.detail()

	fmt.Println("=== CALCULADORA DE PREÇOS ===")
	product1 := factory(Small, 100.0)
	fmt.Println("Preço do produto 1:", product1.CalculatePrice())
	product2 := factory(Medium, 100.0)
	fmt.Println("Preço do produto 2:", product2.CalculatePrice())
	product3 := factory(Large, 100.0)
	fmt.Println("Preço do produto 3:", product3.CalculatePrice())
 
}

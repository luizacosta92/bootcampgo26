package main

import "fmt"

func main() {
	fmt.Println("=== AULA 04 ===")
/*
	var p Product

	produto1 := Product{
		ID:          1,
		Name:        "Aula de Pilates",
		Price:       100.0,
		Description: "Descrição da aula de Pilates",
		Category:    "Categoria 1",
	}

	produto1.Save()
	produto2 := Product{
		ID:          2,
		Name:        "Aula de Funcional",
		Price:       100.0,
		Description: "Descrição da aula de Funcional",
		Category:    "Categoria 2",
	}
	produto2.Save()
	produto3 := Product{
		ID:          3,
		Name:        "Aula de Musculação",
		Price:       100.0,
		Description: "Descrição da aula de Musculação",
		Category:    "Categoria 3",
	}

	produto3.Save()

	p.GetAll()

	fmt.Println("\n🔍 ===== BUSCA POR ID =====")
	var idBuscado int
	fmt.Print("Digite o ID do produto: ")
	fmt.Scanln(&idBuscado)
	
	produto, ok := getById(idBuscado)
	
	if ok {
		fmt.Printf("ID: %d - %s\n", produto.ID, produto.Name)
	} else {
		fmt.Printf("Produto com ID %d não existe\n", idBuscado)
	}*/

	fmt.Println("\n👥 ===== EMPREGADOS =====")
	person1 := Person{
		ID:          1,
		Name:        "João da Silva",
		DateOfBirth: "1990-01-01",
	}
	employee1 := Employee{
		ID:          1,
		Position:    "Gerente",
		Person:      person1,
	}
	employee1.PrintEmployee()
}

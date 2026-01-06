package main

import "fmt"

func main() {
	fmt.Println("=== AULA 04 ===")

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
	
	produto, ok := getById(1)
	fmt.Printf("ID: 1 - %s\n", produto1.Name)
	if ok {
		fmt.Println(produto.Name)
	} else {
		fmt.Println("Produto não encontrado")
	}
}

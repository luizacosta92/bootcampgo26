package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{}

func (p *Product) Save() {
	Products = append(Products, *p)
	fmt.Printf("Produto %s salvo com sucesso\n", p.Name)
}

func (p *Product) GetAll() {
	
	fmt.Println("─────────────────────────────────")
	fmt.Println("Produtos disponíveis:")
	for _, product := range Products {
		fmt.Println(product.Name)
	}

	fmt.Println("─────────────────────────────────")
}

func getById(id int) (Product, bool) {
	for _, product := range Products {
		if product.ID == id {
			return product, true
		}
	}
	return Product{}, false
}

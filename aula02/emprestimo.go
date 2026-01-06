package main

import "fmt"

func Emprestimo(idade int, empregado bool, tempoDeEmpresa int, salario float64) {


	if idade <
	 22 {
		fmt.Println("Cliente deve ter mais de 22 anos")
	} else if !empregado {
		fmt.Println("Cliente deve estar trabalhando")
	} else if tempoDeEmpresa < 1 {
		fmt.Println("Tempo de empresa deve ser maior que 1 ano")
	} else if salario < 100000 {
		fmt.Println("Salário deve ser maior que 100000")
	} else {
		fmt.Println("Emprestimo aprovado")
	}
}

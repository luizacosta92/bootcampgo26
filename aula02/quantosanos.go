package main

import "fmt"

func QuantosAnos() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Nahuel": 26,
		"Brenda": 19,
		"Dario": 44,
		"Pedro": 30,
	}

fmt.Println("Benjamin tem", employees["Benjamin"], "anos")

employees["Federico"] = 25
fmt.Println(employees)

delete(employees, "Pedro")
fmt.Println(employees)

quantidadeDeFuncionarios := len(employees)
fmt.Println("Quantidade de funcionários:", quantidadeDeFuncionarios)

contadorDeFuncionariosMaioresDe21 := 0
for nome, idade := range employees{
	if idade > 21 {
		contadorDeFuncionariosMaioresDe21++
		fmt.Println(nome, "tem", idade, "anos")
	}
}
}
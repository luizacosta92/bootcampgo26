package main

import "math"

func ImpostoSalario(salario float64) float64 {
	var imposto float64
	if salario >= 150000 {
		imposto = salario * 0.27
	} else if salario >= 50000 {
		imposto = salario * 0.17
	} else {
		imposto = salario * 0.08
	}

	return math.Round(imposto*100) / 100
}

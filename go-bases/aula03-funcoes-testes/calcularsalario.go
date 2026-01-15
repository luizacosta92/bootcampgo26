package main

func CalcularSalario(minutosTrabalhados float64, categoria string) (salario float64) {
	switch categoria {
	case "A":
		return minutosTrabalhados / 60 * 3000 * 1.5
	case "B":
		return minutosTrabalhados / 60 * 1500 * 1.2
	case "C":
		return minutosTrabalhados / 60 * 1000
	default:
		return salario
	}
}

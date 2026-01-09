package main

import (
	//"errors"
	"fmt"
)

func main() {
	fmt.Println("=== AULA 03 ===")

	fmt.Println("--- Exercício: Imposto de Salário ---")
	
	var horasTrabalhadas int
	var valorHora float64

	fmt.Print("Digite as horas trabalhadas: ")
	fmt.Scanln(&horasTrabalhadas)
	fmt.Print("Digite o valor da hora: ")
	fmt.Scanln(&valorHora)
	
	salario, err := CalcularSalarioMensal(horasTrabalhadas, valorHora)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	fmt.Println("Salário:", salario)
	
//	salario := 150000.0
//	err := validateSalary(salario)
//
//	if err != nil {
//			fmt.Println(err)
//		} else {
//		fmt.Println("Must pay tax")
//		imposto := ImpostoSalario(salario)
//		fmt.Println("Imposto de salário:", imposto)
//}
//	}
	/*fmt.Println("--- Exercício: Calcular Média ---")
	media, err := CalcularMedia()
	if err != nil {
		fmt.Println("Erro:", err)
	}
	fmt.Println("Média:", media)

	fmt.Println("--- Exercício: Calcular Salário ---")
	salarioVencimento := CalcularSalario(60, "C")
	fmt.Println("Salário:", salarioVencimento)

	fmt.Println("--- Exercício: Calcular Estatísticas ---")
	operacao, err := operation("maximum")
	if err != nil {
		fmt.Println("Erro:", err)
	}
	fmt.Println("Média:", operacao(0))

	fmt.Println("--- Exercício: Calcular Alimentos ---")
	animalDog, _ := calcularAlimentos(dog)
	fmt.Println("Animal:", animalDog(5))
	animalCat, _ := calcularAlimentos(cat)
	fmt.Println("Animal:", animalCat(7))
	animalHamster, _ := calcularAlimentos(hamster)
	fmt.Println("Animal:", animalHamster(9))
	animalTarantula, _ := calcularAlimentos(tarantula)
	fmt.Println("Animal:", animalTarantula(12))*/

}

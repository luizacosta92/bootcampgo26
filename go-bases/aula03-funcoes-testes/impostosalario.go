package main

import (
	"fmt"
	"errors"
	//"math"
)

//type SalaryError struct {
//	Salary int
//}

//	func (e *SalaryError) Error() string {
//		return fmt.Sprintf("Error: salary entered does not reach the taxable minimun")
//	}
var ErrSalaryMinimum = errors.New("Error: salary is less then 100000")

func validateSalary(salario float64) error {
	if salario < 150000 {
		return fmt.Errorf("Error: salary entered is: %.1//f", salario)
	}
	return nil
}

func CalcularSalarioMensal(horasTrabalhadas int, valorHora float64) (float64, error) {
	if horasTrabalhadas < 80 {
		return 0, fmt.Errorf("Error: the worked cannot have worked less than 80 hours per month")
	}
	
	salario := float64(horasTrabalhadas) * valorHora 
	
	if salario >= 150000 {
		imposto := salario - (salario * 0.10)
		salario = salario - imposto
	}
	return salario, nil
}

//type SalaryMinumumError struct{}

//func (e *SalaryMinumumError) Error() string {
//	return "Error: salary is less then 100000"
//}

//var SalaryMinimumError = &SalaryMinumumError{}

//func validateSalary(salario float64) error {
//	if salario < 100000 {
//		return SalaryMinimumError
//	}
//	return nil
//}

//func ImpostoSalario(salario float64) float64 {
//	var imposto float64
//	if salario >= 150000 {
//		imposto = salario * 0.27
//	} else if salario >= 50000 {
//		imposto = salario * 0.17
//	} else {
//		imposto = salario * 0.08
//	}

//	return math.Round(imposto*100) / 100
//}

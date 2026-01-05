package main

import "errors"

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

func operation(tipoCalculo string) (func(...int) float64, error) {

	switch tipoCalculo {
	case minimum:
		return func(notas ...int) float64 {
			if len(notas) == 0 {
				return 0
			}
			min := notas[0]
			for _, nota := range notas {
				if nota < min {
					min = nota
				}
			}
			return float64(min)
		}, nil

	case maximum:
		return func(notas ...int) float64 {
			if len(notas) == 0 {
				return 0
			}
			max := notas[0]
			for _, nota := range notas {
				if nota > max {
					max = nota
				}
			}
			return float64(max)
		}, nil

	case average:
		return func(notas ...int) float64 {
			if len(notas) == 0 {
				return 0
			}
			sum := 0
			for _, nota := range notas {
				sum += nota
			}
			return float64(sum) / float64(len(notas))
		}, nil
	default:
		return nil, errors.New("operação inválida")
	}
}

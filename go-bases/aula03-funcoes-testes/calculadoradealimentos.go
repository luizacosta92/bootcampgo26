package main

import "errors"

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func calcularAlimentos(animal string) (func(quantidade int) float64, error) {
	switch animal {
	case dog:
		return func(quantidade int) float64 {
			return float64(quantidade) * 10
		}, nil
	case cat:
		return func(quantidade int) float64 {
			return float64(quantidade) * 5
		}, nil
	case hamster:
		return func(quantidade int) float64 {
			return float64(quantidade) * 0.25
		}, nil
	case tarantula:
		return func(quantidade int) float64 {
			return float64(quantidade) * 0.15
		}, nil
	default:
		return nil, errors.New("animal inválido")
	}
}

package main

import "errors"

func CalcularMedia(notas ...float64) (float64, error) {
	var media float64

	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("não é possível ter nota negativa")
		}
	}
	for len(notas) == 0 {
		return 0, errors.New("não é possível calcular a média sem notas")
	}
	
	for _, nota := range notas {
		media += nota
	}

	media = media / float64(len(notas))

	return media, nil
}

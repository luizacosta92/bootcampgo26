package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularAlimentos(t *testing.T) {
	t.Run("Dog - 10kg por animal", func(t *testing.T) {
		dogFunc, err := calcularAlimentos("dog")
		assert.NoError(t, err)
		assert.Equal(t, 100.0, dogFunc(10))
	})

	t.Run("Cat - 5kg por animal", func(t *testing.T) {
		catFunc, err := calcularAlimentos("cat")
		assert.NoError(t, err)
		assert.Equal(t, 50.0, catFunc(10))
	})

	t.Run("Hamster - 250g por animal", func(t *testing.T) {
		hamsterFunc, err := calcularAlimentos("hamster")
		assert.NoError(t, err)
		assert.Equal(t, 2.5, hamsterFunc(10))
	})

	t.Run("Tarantula - 150g por animal", func(t *testing.T) {
		tarantulaFunc, err := calcularAlimentos("tarantula")
		assert.NoError(t, err)
		assert.Equal(t, 1.5, tarantulaFunc(10))
	})

	t.Run("Animal inválido - retorna erro", func(t *testing.T) {
		_, err := calcularAlimentos("elefante")
		assert.Error(t, err)
		assert.Equal(t, "animal inválido", err.Error())
	})

	t.Run("Soma total de alimentos", func(t *testing.T) {
		animalDog, _ := calcularAlimentos("dog")
		animalCat, _ := calcularAlimentos("cat")
		animalHamster, _ := calcularAlimentos("hamster")
		animalTarantula, _ := calcularAlimentos("tarantula")

		var amount float64
		amount += animalDog(10)       // 100kg
		amount += animalCat(10)       // 50kg
		amount += animalHamster(10)   // 2.5kg
		amount += animalTarantula(10) // 1.5kg

		assert.Equal(t, 154.0, amount)
	})
}

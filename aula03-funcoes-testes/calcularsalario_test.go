package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularSalarioCategoriaA(t *testing.T) {
	salario := CalcularSalario(60, "A")
	expected := 4500.0
	assert.Equal(t, expected, salario)
}

func TestCalcularSalarioCategoriaB(t *testing.T) {
	salario := CalcularSalario(60, "B")
	expected := 1800.0
	assert.Equal(t, expected, salario)
}

func TestCalcularSalarioCategoriaC(t *testing.T) {
	salario := CalcularSalario(60, "C")
	expected := 1000.0
	assert.Equal(t, expected, salario)
}
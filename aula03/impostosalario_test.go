package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularImpostoSalarioMenorQue50000(t *testing.T) {
	salario := 40000.0

	imposto := ImpostoSalario(salario)

	expected := 3200.0

	assert.Equal(t, expected, imposto)
}

func TestCalcularImpostoSalarioMaiorQue50000(t *testing.T) {
	salario := 50001.0

	imposto := ImpostoSalario(salario)

	expected := 8500.17

	assert.Equal(t, expected, imposto, 0.01)
}

func TestCalcularImpostoSalarioMaiorQue150000(t *testing.T) {
	salario := 150001.0

	imposto := ImpostoSalario(salario)

	expected := 40500.27

	assert.Equal(t, expected, imposto, 0.01)
}

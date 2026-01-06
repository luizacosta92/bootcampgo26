package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcularAlimentosCachorro(t *testing.T) {
	animal, err := calcularAlimentos("dog")
	require.NoError(t, err)
	require.Equal(t, 100.0, animal(10))
}

func TestCalcularAlimentosGato(t *testing.T) {
	animal, err := calcularAlimentos("cat")
	require.NoError(t, err)
	require.Equal(t, 25.0, animal(5))
}

func TestCalcularAlimentosHamster(t *testing.T) {
	animal, err := calcularAlimentos("hamster")
	require.NoError(t, err)
	require.Equal(t, 2.5, animal(10))
}

func TestCalcularAlimentosTarantula(t *testing.T) {
	animal, err := calcularAlimentos("tarantula")
	expected := 1.5
	require.NoError(t, err)
	require.Equal(t, expected, animal(10))
}

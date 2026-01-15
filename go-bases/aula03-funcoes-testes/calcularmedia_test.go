package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularMedia(t *testing.T) {
	notas := []float64{10, 10, 10}

	media, err := CalcularMedia(notas...)

	assert.Equal(t, 10.0, media)
	assert.NoError(t, err)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOperationMinimum(t *testing.T) {

	minimum, err := operation("minimum")
	require.NoError(t, err)
	require.Equal(t, 2.0, minimum(5, 8, 2, 10))
}

func TestOperationMaximum(t *testing.T) {
	maximum, err := operation("maximum")
	require.NoError(t, err)
	require.Equal(t, 10.0, maximum(5, 8, 2, 10))
}

func TestOperationAverage(t *testing.T) {
	average, err := operation("average")
	require.NoError(t, err)
	require.Equal(t, 6.25, average(5, 8, 2, 10))
}


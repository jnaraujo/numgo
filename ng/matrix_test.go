package ng

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMatrix(t *testing.T) {
	expected := []Array[float64]{{1, 2.5, 3}, {4, 5, 6}}
	got := NewMatrix(NewArray(1, 2.5, 3), NewArray[float64](4, 5, 6))
	assert.Equal(t, expected, []Array[float64](got))
}

func TestZeroesMatrix(t *testing.T) {
	expected := []Array[int]{{0, 0, 0}, {0, 0, 0}}
	got := ZeroesMatrix[int](2, 3)
	assert.Equal(t, expected, []Array[int](got))
}

func TestOnesMatrix(t *testing.T) {
	expected := []Array[float64]{{1, 1, 1}, {1, 1, 1}}
	got := OnesMatrix[float64](2, 3)
	assert.Equal(t, expected, []Array[float64](got))
}

func TestMatrixString(t *testing.T) {
	expected := "1, 2.5, 3\n4, 5, 6"
	got := NewMatrix(NewArray(1, 2.5, 3), NewArray[float64](4, 5, 6)).String()
	assert.Equal(t, expected, got)
}

func TestMatrixOnesString(t *testing.T) {
	expected := "1, 1, 1\n1, 1, 1"
	got := OnesMatrix[int](2, 3).String()
	assert.Equal(t, expected, got)
}

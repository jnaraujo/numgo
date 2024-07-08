package ng

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateIntArrayWithNewArray(t *testing.T) {
	expected := []int{1, 2, -3}
	got := NewArray(1, 2, -3)
	assert.Equal(t, expected, []int(got))
}

func TestCreateEmptyArrayWithNewArray(t *testing.T) {
	expected := []int{}
	got := NewArray[int]()
	assert.Equal(t, expected, []int(got))
}

func TestCreateIntArray(t *testing.T) {
	expected := []int{1, 2, -3}
	got := Array[int]{1, 2, -3}
	assert.Equal(t, expected, []int(got))
}

func TestCreateFloatArray(t *testing.T) {
	expected := []float32{9999, 21312, -23431234, -3.10}
	got := Array[float32]{9999, 21312, -23431234, -3.10}
	assert.Equal(t, expected, []float32(got))
}

func TestArrayMutability(t *testing.T) {
	expected := []int{120, 2, -3}
	got := Array[int]{1, 2, -3}
	got[0] = 120
	assert.Equal(t, expected, []int(got))
}

func TestCreateIntArrayWithZeroes(t *testing.T) {
	expected := []int{0, 0, 0}
	got := Zeroes[int](3)
	assert.Equal(t, expected, []int(got))
}

func TestCreateFloatArrayWithOnes(t *testing.T) {
	expected := []float32{1, 1, 1}
	got := Ones[float32](3)
	assert.Equal(t, expected, []float32(got))
}

func TestRangeArray(t *testing.T) {
	expected := []int{0, 2, 4}
	got := Range(0, 6, 2)
	assert.Equal(t, expected, []int(got))
}

func TestRangeFloatArray(t *testing.T) {
	expected := []float64{0, 0.5, 1, 1.5}
	got := Range(0, 2, 0.5)
	assert.Equal(t, expected, []float64(got))
}

func TestRangeFloatArrayWithNegativeRange(t *testing.T) {
	expected := []int{-2, -1, 0, 1}
	got := Range(-2, 2, 1)
	assert.Equal(t, expected, []int(got))
}

func TestArraySort(t *testing.T) {
	expected := NewArray(-2, -1.5, 0, 1, 99.99)
	arr := NewArray(1, -2, 0, 99.99, -1.5)
	assert.NotEqual(t, arr, expected)
	assert.Equal(t, expected, arr.Sort())
}

func TestArrayConcatenation(t *testing.T) {
	expected := NewArray(1, 2, 3, 4, 5, 6)
	arr1 := NewArray(1, 2, 3)
	assert.Equal(t, expected, arr1.Concat(NewArray(4, 5, 6)))
}

func TestArrayConcatenationWithDifferentTypes(t *testing.T) {
	expected := NewArray(-10, 2.2, 3.5, 4, 5, 6)
	arr1 := NewArray(-10, 2.2, 3.5)
	assert.Equal(t, expected, arr1.Concat(NewArray[float64](4, 5, 6)))
}
func TestArraySum(t *testing.T) {
	expected := NewArray(-10, 5.5, 20)
	arr1 := NewArray[float64](10, 2, 10)
	assert.Equal(t, expected, arr1.Sum(NewArray(-20, 3.5, 10)))
}

func TestArrayMultiplyScalar(t *testing.T) {
	expected := NewArray(-35, 19.25, 70)
	arr := NewArray(-10, 5.5, 20)
	assert.Equal(t, expected, arr.MultiplyScalar(3.5))
}

func TestArrayMultiply(t *testing.T) {
	expected := NewArray[float64](-200, 11, 100)
	arr1 := NewArray(-10, 5.5, 20)
	assert.Equal(t, expected, arr1.Multiply(NewArray[float64](20, 2, 5)))
}

func TestArrayString(t *testing.T) {
	expected := "-10, 5.5, 20"
	arr1 := NewArray(-10, 5.5, 20)
	assert.Equal(t, expected, arr1.String())
}

package ng

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/jnaraujo/numgo/internal"
)

type Array[T Scalar] []T

func NewArray[T Scalar](elements ...T) Array[T] {
	if len(elements) == 0 {
		return Array[T]{}
	}
	return Array[T](elements)
}

func Zeroes[T Scalar](n int) Array[T] {
	return make(Array[T], n)
}

// Ones creates an array of n elements, all set to 1.
func Ones[T Scalar](n int) Array[T] {
	a := make(Array[T], n)
	for i := range a {
		a[i] = 1
	}
	return a
}

// create array with a range of Scalars
func Range[T Scalar](start, end, step T) Array[T] {
	internal.Assert(step != 0, "step cannot be zero")
	internal.Assert(end != start, "end cannot be equal to start")
	internal.Assert(end > start, "end must be greater than start")

	arr := make(Array[T], int((end-start)/step))
	for i := 0; i < len(arr); i++ {
		arr[i] = start + T(i)*step
	}

	return arr
}

func (arr Array[T]) Sort() Array[T] {
	arr = slices.Clone(arr)
	slices.Sort(arr)
	return arr
}

func (arr Array[T]) Concat(other Array[T]) Array[T] {
	return append(arr, other...)
}

func (arr Array[T]) Sum(other Array[T]) Array[T] {
	internal.Assert(len(arr) == len(other), "arrays must have the same length")
	array := make(Array[T], len(arr))
	for i := 0; i < len(array); i++ {
		array[i] = arr[i] + other[i]
	}
	return array
}

func (arr Array[T]) Sub(other Array[T]) Array[T] {
	internal.Assert(len(arr) == len(other), "arrays must have the same length")
	array := make(Array[T], len(arr))
	for i := 0; i < len(array); i++ {
		array[i] = arr[i] - other[i]
	}
	return array
}

func (arr Array[T]) Power(n float64) Array[T] {
	array := make(Array[T], len(arr))
	for i := 0; i < len(array); i++ {
		array[i] = T(math.Pow(float64(arr[i]), n))
	}
	return array
}

func (arr Array[T]) Multiply(other Array[T]) Array[T] {
	internal.Assert(len(arr) == len(other), "arrays must have the same length")
	array := make(Array[T], len(arr))
	for i := 0; i < len(array); i++ {
		array[i] = arr[i] * other[i]
	}
	return array
}

func (arr Array[T]) MultiplyScalar(scalar T) Array[T] {
	arr = slices.Clone(arr)
	for i := range arr {
		arr[i] *= scalar
	}
	return arr
}

func (arr Array[T]) Dot(other Array[T]) T {
	internal.Assert(len(arr) == len(other), "arrays must have the same length")
	sum := T(0)
	for i := 0; i < len(arr); i++ {
		sum += arr[i] * other[i]
	}
	return sum
}

func (arr Array[T]) String() string {
	var s strings.Builder
	for i, el := range arr {
		s.WriteString(fmt.Sprintf("%v", el))
		if i != len(arr)-1 {
			s.WriteString(", ")
		}
	}
	return s.String()
}

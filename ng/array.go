package ng

import (
	"numgo/internal"
	"slices"
)

type Array[T number] []T

func NewArray[T number](elements ...T) Array[T] {
	if len(elements) == 0 {
		return Array[T]{}
	}
	return Array[T](elements)
}

func Zeroes[T number](n int) Array[T] {
	return make(Array[T], n)
}

// Ones creates an array of n elements, all set to 1.
func Ones[T number](n int) Array[T] {
	a := make(Array[T], n)
	for i := range a {
		a[i] = 1
	}
	return a
}

// create array with a range of numbers
func Range[T number](start, end, step T) Array[T] {
	internal.Assert(step != 0, "step cannot be zero")
	internal.Assert(end != start, "end cannot be equal to start")
	internal.Assert(end > start, "end must be greater than start")

	arr := make(Array[T], int((end-start)/step))
	for i := 0; i < len(arr); i++ {
		arr[i] = start + T(i)*step
	}

	return arr
}

func Sort[T number](arr Array[T]) Array[T] {
	arr = slices.Clone(arr)
	slices.Sort(arr)
	return arr
}

func Concatenate[T number](arr1, arr2 Array[T]) Array[T] {
	return append(arr1, arr2...)
}

func Sum[T number](arr1, arr2 Array[T]) Array[T] {
	internal.Assert(len(arr1) == len(arr2), "arrays must have the same length")
	arr := make(Array[T], len(arr1))
	for i := 0; i < len(arr); i++ {
		arr[i] = arr1[i] + arr2[i]
	}
	return arr
}

func Multiply[T number](arr1, arr2 Array[T]) Array[T] {
	internal.Assert(len(arr1) == len(arr2), "arrays must have the same length")
	arr := make(Array[T], len(arr1))
	for i := 0; i < len(arr); i++ {
		arr[i] = arr1[i] * arr2[i]
	}
	return arr
}

func MultiplyScalar[T number](arr Array[T], scalar T) Array[T] {
	arr = slices.Clone(arr)
	for i := range arr {
		arr[i] *= scalar
	}
	return arr
}

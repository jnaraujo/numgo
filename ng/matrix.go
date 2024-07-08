package ng

import (
	"fmt"
	"strings"
)

type Matrix[T Scalar] []Array[T]

func NewMatrix[T Scalar](elements ...Array[T]) Matrix[T] {
	if len(elements) == 0 {
		return Matrix[T]{}
	}
	return elements
}

func ZeroesMatrix[T Scalar](m, n int) Matrix[T] {
	mat := make(Matrix[T], m)
	for i := 0; i < m; i++ {
		mat[i] = Zeroes[T](n)
	}
	return mat
}

func OnesMatrix[T Scalar](m, n int) Matrix[T] {
	mat := make(Matrix[T], m)
	for i := 0; i < m; i++ {
		mat[i] = Ones[T](n)
	}
	return mat
}

func (mat Matrix[T]) String() string {
	var s strings.Builder
	for i, arr := range mat {
		s.WriteString(fmt.Sprintf("%v", arr))
		if i < len(mat)-1 {
			s.WriteString("\n")
		}
	}
	return s.String()
}

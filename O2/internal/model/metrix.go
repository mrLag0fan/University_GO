package model

import (
	"errors"
)

type Matrix[T int64 | float64] struct {
	n      int
	m      int
	matrix [][]T
}

func NewZeroMatrix[T int64 | float64](n int, m int) *Matrix[T] {
	matrix := make([][]T, n)
	for i := range matrix {
		matrix[i] = make([]T, m)
	}
	return &Matrix[T]{
		n:      n,
		m:      m,
		matrix: matrix,
	}
}

func NewMatrix[T int64 | float64](data [][]T) *Matrix[T] {
	matrix := make([][]T, len(data))
	for i := range matrix {
		matrix[i] = make([]T, len(data[i]))
		for j := range matrix[i] {
			matrix[i][j] = data[i][j]
		}
	}
	return &Matrix[T]{
		n:      len(data),
		m:      len(data[0]),
		matrix: matrix,
	}
}

func (matrix *Matrix[T]) Set(i, j int, data T) error {
	if (i >= 0 && i < len(matrix.matrix)) && (j >= 0 && j < len(matrix.matrix[i])) {
		matrix.matrix[i][j] = data
		return nil
	}
	return errors.New("index out of bounds")
}

func (matrix *Matrix[T]) Get(i, j int) (T, error) {
	if (i >= 0 && i < len(matrix.matrix)) && (j >= 0 && j < len(matrix.matrix[i])) {
		return matrix.matrix[i][j], nil
	}
	return -1, errors.New("index out of bounds")
}

func (matrix *Matrix[T]) Add(otherMatrix Matrix[T]) *Matrix[T] {
	normalize(matrix, &otherMatrix)
	res := NewZeroMatrix[T](len(matrix.matrix), len(matrix.matrix[0]))
	for i, row := range res.matrix {
		for j, _ := range row {
			res.matrix[i][j] = matrix.matrix[i][j] + otherMatrix.matrix[i][j]
		}
	}
	return res
}

func normalize[T int64 | float64](matrix, otherMatrix *Matrix[T]) {
	matrix.makeRectangularMatrix()
	otherMatrix.makeRectangularMatrix()
	if len(matrix.matrix[0]) > len(otherMatrix.matrix[0]) {
		otherMatrix.adjustToSize(len(matrix.matrix[0]))
	} else {
		matrix.adjustToSize(len(otherMatrix.matrix[0]))
	}
}

func (matrix *Matrix[T]) adjustToSize(newRowSize int) {
	additionalLength := newRowSize - len(matrix.matrix[0])
	if additionalLength <= 0 {
		return
	}
	for i := range matrix.matrix {
		for j := 0; j < additionalLength; j++ {
			matrix.matrix[i] = append(matrix.matrix[i], 0)
		}
	}
}

func (matrix *Matrix[T]) isJaggedMatrix() bool {
	for _, row := range matrix.matrix {
		if len(row) != matrix.m {
			return false
		}
	}
	return true
}

func (matrix *Matrix[T]) makeRectangularMatrix() {
	if matrix.isJaggedMatrix() {
		for i := range matrix.matrix {
			additionalLength := matrix.maxRowSize() - len(matrix.matrix[i])
			if additionalLength <= 0 {
				continue
			}
			for j := 0; j < additionalLength; j++ {
				matrix.matrix[i] = append(matrix.matrix[i], 0)
			}
		}
	}
}

func (matrix *Matrix[T]) maxRowSize() int {
	maxLength := len(matrix.matrix[0])
	for _, row := range matrix.matrix {
		if len(row) > maxLength {
			maxLength = len(row)
		}
	}
	return maxLength
}

package model

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

func (matrix *Matrix[T]) Print() {
	for _, row := range matrix.matrix {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

func (matrix *Matrix[T]) Input() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введіть розмір матриці (рядки стовпці):")
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			fmt.Println("Неправильний формат вводу")
			return
		}
		n, err1 := strconv.Atoi(parts[0])
		m, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Неправильний формат вводу")
			return
		}
		matrix.n = n
		matrix.m = m
		matrix.matrix = make([][]T, n)
		fmt.Println("Введіть елементи матриці:")
		for i := 0; i < n; i++ {
			if scanner.Scan() {
				line := scanner.Text()
				parts := strings.Split(line, " ")
				if len(parts) != m {
					fmt.Println("Неправильна кількість елементів у рядку")
					return
				}
				matrix.matrix[i] = make([]T, m)
				for j := 0; j < m; j++ {
					val, err := strconv.ParseInt(parts[j], 10, 64)
					if err != nil {
						fmt.Println("Неправильний формат елементу")
						return
					}
					matrix.matrix[i][j] = T(val)
				}
			}
		}
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
	if matrix.n == otherMatrix.n && matrix.m == otherMatrix.m {
		res := NewZeroMatrix[T](len(matrix.matrix), len(matrix.matrix[0]))
		for i, row := range res.matrix {
			for j := range row {
				res.matrix[i][j] = matrix.matrix[i][j] + otherMatrix.matrix[i][j]
			}
		}
		return res
	}
	return nil
}

func (matrix *Matrix[T]) Subtract(otherMatrix Matrix[T]) *Matrix[T] {
	if matrix.n == otherMatrix.n && matrix.m == otherMatrix.m {
		res := NewZeroMatrix[T](len(matrix.matrix), len(matrix.matrix[0]))
		for i, row := range res.matrix {
			for j := range row {
				res.matrix[i][j] = matrix.matrix[i][j] - otherMatrix.matrix[i][j]
			}
		}
		return res
	}
	return nil
}

func (matrix *Matrix[T]) Multiple(otherMatrix Matrix[T]) *Matrix[T] {
	if matrix.m == otherMatrix.n {
		res := NewZeroMatrix[T](matrix.n, otherMatrix.m)
		for i := 0; i < matrix.n; i++ {
			for j := 0; j < otherMatrix.m; j++ {
				for k := 0; k < matrix.m; k++ {
					res.matrix[i][j] += matrix.matrix[i][k] * otherMatrix.matrix[k][j]
				}
			}
		}
		return res
	}
	return nil
}

func (matrix *Matrix[T]) Determinant() float64 {
	if matrix.n == matrix.m {
		res, err := matrix.elimination()
		if err != nil {
			return 0
		}
		var det float64 = 1
		for i := 0; i < matrix.n; i++ {
			det = float64(res.matrix[i][i]) * det
		}
		return det
	}
	return 0
}

func (matrix *Matrix[T]) SolveSLAR() ([]float64, error) {
	if matrix.n == matrix.m-1 {
		triangle, err := matrix.elimination()
		if err != nil {
			return []float64{}, errors.New("Infinite number of solutions")
		}
		results := make([]float64, matrix.n)
		for i := triangle.n - 1; i >= 0; i-- {
			results[i] = float64(triangle.matrix[i][triangle.n])
			for j := i + 1; j < triangle.n; j++ {
				results[i] -= float64(triangle.matrix[i][j]) * results[j]
			}
			results[i] /= float64(triangle.matrix[i][i])
		}
		return results, nil
	}
	return nil, errors.New("No solutions")
}

func (matrix *Matrix[T]) Transpose() *Matrix[T] {
	res := NewZeroMatrix[T](matrix.m, matrix.n)
	for i, row := range matrix.matrix {
		for j := range row {
			res.matrix[j][i] = matrix.matrix[i][j]
		}
	}
	return res
}

func (matrix *Matrix[T]) Inverse() (*Matrix[float64], error) {
	if matrix.n == matrix.m {
		det := matrix.Determinant()
		if det == 0 {
			return nil, errors.New("Матриця є невиродженою, обернена матриця не існує")
		}
		res := NewZeroMatrix[float64](matrix.m, matrix.n)
		for i, row := range res.matrix {
			for j := range row {
				res.matrix[i][j] = matrix.cofactor(i, j) / det
			}
		}

		return res, nil
	}
	return nil, errors.New("Матриця не є квадратною")
}

func (matrix *Matrix[T]) cofactor(row, col int) float64 {
	minorDet := matrix.minor(row, col).Determinant()
	return math.Pow(-1, float64(row+col)) * minorDet
}

func (matrix *Matrix[T]) minor(row, col int) *Matrix[T] {
	minor := NewZeroMatrix[T](matrix.n-1, matrix.m-1)

	for i, r := range matrix.matrix {
		if i == row {
			continue
		}
		for j, val := range r {
			if j == col {
				continue
			}
			minorRow := i
			minorCol := j
			if i > row {
				minorRow = i - 1
			}
			if j > col {
				minorCol = j - 1
			}
			minor.matrix[minorRow][minorCol] = val
		}
	}

	return minor
}

func (matrix *Matrix[T]) elimination() (*Matrix[float64], error) {
	res := NewZeroMatrix[float64](matrix.n, matrix.m)
	for i := 0; i < matrix.n; i++ {
		for j := 0; j < matrix.m; j++ {
			res.matrix[i][j] = float64(matrix.matrix[i][j])
		}
	}

	for i := 0; i < matrix.n; i++ {
		for j := i + 1; j < matrix.n; j++ {
			if res.matrix[i][i] == 0 {
				return nil, errors.New("Division by zero")
			}
			factor := res.matrix[j][i] / res.matrix[i][i]
			for k := i; k < matrix.n; k++ {
				res.matrix[j][k] -= factor * res.matrix[i][k]
			}
		}
	}
	return res, nil
}

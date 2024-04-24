package main

import (
	"fmt"
	"math"
	"sync"
)

func gaussJordanParallel(a [][]float64, b []float64) ([]float64, error) {
	n := len(a)
	if n == 0 || len(a[0]) != n || len(b) != n {
		return nil, fmt.Errorf("incorrect input: len(a)=%d, len(a[0])=%d, len(b)=%d", len(a), len(a[0]), len(b))
	}

	// Initialize the result vector x with the values of b
	x := make([]float64, n)
	copy(x, b)

	// Perform the Gauss-Jordan elimination in parallel
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Find pivot row and swap
			max := math.Abs(a[i][i])
			pivot := i
			for j := i + 1; j < n; j++ {
				absVal := math.Abs(a[j][i])
				if absVal > max {
					max = absVal
					pivot = j
				}
			}
			a[i], a[pivot] = a[pivot], a[i]
			x[i], x[pivot] = x[pivot], x[i]

			// Eliminate the current column
			for j := 0; j < n; j++ {
				if i != j {
					ratio := a[j][i] / a[i][i]
					for k := i; k < n; k++ {
						a[j][k] -= ratio * a[i][k]
					}
					x[j] -= ratio * x[i]
				}
			}

			// Normalize the current row
			div := a[i][i]
			for j := i; j < n; j++ {
				a[i][j] /= div
			}
			x[i] /= div
		}(i)
	}
	wg.Wait()

	return x, nil
}

func main() {
	// Example input
	a := [][]float64{
		{2, 1, -1},
		{-3, -1, 2},
		{-2, 1, 2},
	}
	b := []float64{8, -11, -3}

	// Solve the system of linear equations
	x, err := gaussJordanParallel(a, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Output the solution
	fmt.Println("Solution:")
	for i, val := range x {
		fmt.Printf("x%d = %f\n", i+1, val)
	}
}

package model

import (
	"reflect"
	"testing"
)

func TestNewZeroMatrixInt64(t *testing.T) {
	type args struct {
		n int
		m int
	}

	type testCase[T interface{ int64 | float64 }] struct {
		name string
		args args
		want *Matrix[T]
	}
	tests := []testCase[int64]{
		{
			name: "3x2 Matrix Int64",
			args: args{n: 3, m: 2},
			want: &Matrix[int64]{
				n: 3,
				m: 2,
				matrix: [][]int64{
					{0, 0},
					{0, 0},
					{0, 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewZeroMatrix[int64](tt.args.n, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewZeroMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewZeroMatrixFloat64(t *testing.T) {
	type args struct {
		n int
		m int
	}

	type testCase[T interface{ int64 | float64 }] struct {
		name string
		args args
		want *Matrix[T]
	}

	tests := []testCase[float64]{
		{
			name: "2x3 Matrix Float64",
			args: args{n: 2, m: 3},
			want: &Matrix[float64]{
				n: 2,
				m: 3,
				matrix: [][]float64{
					{0, 0, 0},
					{0, 0, 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewZeroMatrix[float64](tt.args.n, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewZeroMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Set(t *testing.T) {
	type args[T interface{ int64 | float64 }] struct {
		i    int
		j    int
		data T
	}
	type testCase[T interface{ int64 | float64 }] struct {
		name    string
		matrix  Matrix[T]
		args    args[T]
		wantErr bool
	}
	tests := []testCase[int64]{
		{
			name: "Set valid element",
			matrix: Matrix[int64]{
				n: 2,
				m: 3,
				matrix: [][]int64{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			args:    args[int64]{i: 0, j: 0, data: 10},
			wantErr: false,
		},
		{
			name: "Set invalid index",
			matrix: Matrix[int64]{
				n: 2,
				m: 3,
				matrix: [][]int64{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			args:    args[int64]{i: 2, j: 2, data: 30},
			wantErr: true,
		},
		{
			name: "Set invalid negative index",
			matrix: Matrix[int64]{
				n: 2,
				m: 3,
				matrix: [][]int64{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			args:    args[int64]{i: -1, j: 0, data: 30},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.matrix.Set(tt.args.i, tt.args.j, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Перевірка, що значення в матриці встановлене правильно
			if !tt.wantErr && tt.matrix.matrix[tt.args.i][tt.args.j] != tt.args.data {
				t.Errorf("Set() failed. Expected value: %v, Got: %v", tt.args.data, tt.matrix.matrix[tt.args.i][tt.args.j])
			}
		})
	}
}

func TestMatrix_Get(t *testing.T) {
	type args struct {
		i int
		j int
	}
	type testCase[T interface{ int64 | float64 }] struct {
		name    string
		matrix  Matrix[T]
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[int64]{
		{
			name: "Get valid element",
			matrix: Matrix[int64]{
				n: 2,
				m: 3,
				matrix: [][]int64{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			args:    args{i: 1, j: 2},
			want:    6,
			wantErr: false,
		},
		{
			name: "Get invalid index",
			matrix: Matrix[int64]{
				n: 2,
				m: 3,
				matrix: [][]int64{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			args:    args{i: 2, j: 2},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.matrix.Get(tt.args.i, tt.args.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Add(t *testing.T) {
	type args[T interface{ int64 | float64 }] struct {
		otherMatrix Matrix[int64]
	}
	type testCase[T interface{ int64 | float64 }] struct {
		name   string
		matrix Matrix[T]
		args   args[T]
		want   *Matrix[T]
	}
	tests := []testCase[int64]{
		{
			name: "Add matrices of same size",
			matrix: Matrix[int64]{
				n: 2,
				m: 3,
				matrix: [][]int64{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			args: args[int64]{
				otherMatrix: Matrix[int64]{
					n: 2,
					m: 3,
					matrix: [][]int64{
						{7, 8, 9},
						{10, 11, 12},
					},
				},
			},
			want: &Matrix[int64]{
				n: 2,
				m: 3,
				matrix: [][]int64{
					{8, 10, 12},
					{14, 16, 18},
				},
			},
		},
		{
			name: "Add matrices of different sizes",
			matrix: Matrix[int64]{
				n: 2,
				m: 3,
				matrix: [][]int64{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			args: args[int64]{
				otherMatrix: Matrix[int64]{
					n: 3,
					m: 2,
					matrix: [][]int64{
						{7, 8},
						{9, 10},
						{11, 12},
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.Add(tt.args.otherMatrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Subtract(t *testing.T) {
	type args[T interface{ int64 | float64 }] struct {
		otherMatrix Matrix[T]
	}
	type testCase[T interface{ int64 | float64 }] struct {
		name   string
		matrix Matrix[T]
		args   args[T]
		want   *Matrix[T]
	}
	tests := []testCase[int64]{
		{
			name: "Subtract matrices of the same size",
			matrix: Matrix[int64]{
				n: 2,
				m: 3,
				matrix: [][]int64{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			args: args[int64]{
				otherMatrix: Matrix[int64]{
					n: 2,
					m: 3,
					matrix: [][]int64{
						{7, 8, 9},
						{10, 11, 12},
					},
				},
			},
			want: &Matrix[int64]{
				n: 2,
				m: 3,
				matrix: [][]int64{
					{-6, -6, -6},
					{-6, -6, -6},
				},
			},
		},
		{
			name: "Subtract matrices of different sizes",
			matrix: Matrix[int64]{
				n: 2,
				m: 3,
				matrix: [][]int64{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			args: args[int64]{
				otherMatrix: Matrix[int64]{
					n: 3,
					m: 2,
					matrix: [][]int64{
						{7, 8},
						{9, 10},
						{11, 12},
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.Subtract(tt.args.otherMatrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Multiple(t *testing.T) {
	type args[T interface{ int64 | float64 }] struct {
		otherMatrix Matrix[T]
	}
	type testCase[T interface{ int64 | float64 }] struct {
		name   string
		matrix Matrix[T]
		args   args[T]
		want   *Matrix[T]
	}
	tests := []testCase[int64]{
		{
			name: "Multiply matrices of compatible sizes",
			matrix: Matrix[int64]{
				n: 2,
				m: 3,
				matrix: [][]int64{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			args: args[int64]{
				otherMatrix: Matrix[int64]{
					n: 3,
					m: 2,
					matrix: [][]int64{
						{7, 8},
						{9, 10},
						{11, 12},
					},
				},
			},
			want: &Matrix[int64]{
				n: 2,
				m: 2,
				matrix: [][]int64{
					{58, 64},
					{139, 154},
				},
			},
		},
		{
			name: "Multiply matrices of incompatible sizes",
			matrix: Matrix[int64]{
				n: 2,
				m: 2,
				matrix: [][]int64{
					{1, 2},
					{4, 5},
				},
			},
			args: args[int64]{
				otherMatrix: Matrix[int64]{
					n: 3,
					m: 2,
					matrix: [][]int64{
						{7, 8},
						{9, 10},
						{11, 12},
					},
				},
			},
			want: nil, // Change this to the expected result for matrices of incompatible sizes
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.Multiple(tt.args.otherMatrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Determinant(t *testing.T) {
	type testCase[T interface{ int64 | float64 }] struct {
		name   string
		matrix Matrix[T]
		want   float64
	}
	tests := []testCase[int64]{
		{
			name: "Determinant of a 2x2 matrix",
			matrix: Matrix[int64]{
				n: 2,
				m: 2,
				matrix: [][]int64{
					{1, 2},
					{3, 4},
				},
			},
			want: -2,
		},
		{
			name: "Determinant of a 3x3 matrix",
			matrix: Matrix[int64]{
				n: 3,
				m: 3,
				matrix: [][]int64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: 0,
		},
		{
			name: "Determinant of a 2x2 matrix with positive determinant",
			matrix: Matrix[int64]{
				n: 2,
				m: 2,
				matrix: [][]int64{
					{2, 3},
					{1, 4},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.Determinant(); got != tt.want {
				t.Errorf("Determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_elimination(t *testing.T) {
	type testCase[T interface{ int64 | float64 }] struct {
		name   string
		matrix Matrix[T]
		want   *Matrix[float64]
	}
	tests := []testCase[float64]{
		{
			name: "Elimination of a 3x3 matrix",
			matrix: Matrix[float64]{
				n: 3,
				m: 3,
				matrix: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: &Matrix[float64]{
				n: 3,
				m: 3,
				matrix: [][]float64{
					{1, 2, 3},
					{0, -3, -6},
					{0, 0, 0},
				},
			},
		},
		{
			name: "Elimination of a 2x2 matrix",
			matrix: Matrix[float64]{
				n: 2,
				m: 2,
				matrix: [][]float64{
					{2, 3},
					{1, 4},
				},
			},
			want: &Matrix[float64]{
				n: 2,
				m: 2,
				matrix: [][]float64{
					{2, 3},
					{0, 2.5},
				},
			},
		},
		{
			name: "Elimination of a 4x4 matrix",
			matrix: Matrix[float64]{
				n: 4,
				m: 4,
				matrix: [][]float64{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 16},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.matrix.elimination(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("elimination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_SolveSLAR(t *testing.T) {
	type testCase[T interface{ int64 | float64 }] struct {
		name    string
		matrix  Matrix[T]
		want    []float64
		wantErr bool
	}
	tests := []testCase[float64]{
		{
			name: "3x3 System",
			matrix: Matrix[float64]{
				n: 3,
				m: 4,
				matrix: [][]float64{
					{1, 1, 1, 6},
					{0, 2, 5, 8},
					{2, 5, -1, 9},
				},
			},
			want:    []float64{5, 1, -2},
			wantErr: false,
		},
		{
			name: "2x2 System",
			matrix: Matrix[float64]{
				n: 2,
				m: 3,
				matrix: [][]float64{
					{2, -1, 8},
					{-3, 1, -4},
				},
			},
			want:    []float64{-4, -16},
			wantErr: false,
		},
		// Add more test cases as needed...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.matrix.SolveSLAR()
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveSLAR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveSLAR() got = %v, want %v", got, tt.want)
			}
		})
	}
}

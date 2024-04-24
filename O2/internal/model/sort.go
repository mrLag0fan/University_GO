package model

type RowSortedMatrix[T int64 | float64] struct {
	*Matrix[T]
}

func (rsm *RowSortedMatrix[T]) Less(i, j int) bool {
	for k := 0; k < rsm.m; k++ {
		if rsm.Matrix.matrix[i][k] != rsm.Matrix.matrix[j][k] {
			return rsm.Matrix.matrix[i][k] < rsm.Matrix.matrix[j][k]
		}
	}
	return false
}

func (rsm *RowSortedMatrix[T]) Swap(i, j int) {
	rsm.Matrix.matrix[i], rsm.Matrix.matrix[j] = rsm.Matrix.matrix[j], rsm.Matrix.matrix[i]
}

func (rsm *RowSortedMatrix[T]) Len() int {
	return rsm.n
}

type ColSortedMatrix[T int64 | float64] struct {
	*Matrix[T]
}

func (csm *ColSortedMatrix[T]) Less(i, j int) bool {
	for k := 0; k < csm.n; k++ {
		if csm.Matrix.matrix[k][i] != csm.Matrix.matrix[k][j] {
			return csm.Matrix.matrix[k][i] < csm.Matrix.matrix[k][j]
		}
	}
	return false
}

func (csm *ColSortedMatrix[T]) Swap(i, j int) {
	for k := 0; k < csm.n; k++ {
		csm.Matrix.matrix[k][i], csm.Matrix.matrix[k][j] = csm.Matrix.matrix[k][j], csm.Matrix.matrix[k][i]
	}
}

func (csm *ColSortedMatrix[T]) Len() int {
	return csm.m
}

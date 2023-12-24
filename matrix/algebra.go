package matrix

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var ErrDivByZero = errors.New("divide by zero")
var ErrMatSizeMissmatch = errors.New("matrix size missmatch")

// MatrixAdd adds two 2D matrices of the same size
// Input:
// a, b are two matrices of the form [rows][column]Matrix
// Output:
// resVal is the sum matrix
func MatrixAdd[Ord constraints.Ordered](a, b [][]Ord) (resVal [][]Ord, err error) {
	sizeA := MatrixSize(a)
	sizeB := MatrixSize(b)
	if sizeA != sizeB {
		return nil, ErrMatSizeMissmatch
	}
	rows := sizeA[0]
	columns := sizeA[1]
	// Expand output slice
	for i := 0; i < rows; i++ {
		resVal = append(resVal, make([]Ord, columns))
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			resVal[i][j] = a[i][j] + b[i][j]
		}
	}
	return resVal, nil
}

// MatrixSize estimates the size of a 2D matrix
// Input:
// input is the numerical input matrix of the form [rows][column]Matrix
// Output:
// matrixSize is a vector with the size of the matrix [rows, columns]
func MatrixSize[Ord constraints.Ordered](input [][]Ord) (matrixSize [2]int) {
	matrixSize = [2]int{len(input[:][:]), len(input[0][:])}
	return
}

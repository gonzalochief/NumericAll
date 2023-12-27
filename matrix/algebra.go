package matrix

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var ErrDivByZero = errors.New("divide by zero")
var ErrMatSizeMissmatch = errors.New("matrix size missmatch")

type Number interface {
	constraints.Float | constraints.Integer | constraints.Complex
}

// MatrixAdd adds two 2D matrices of the same size
// Input:
// a, b are two matrices of the form [rows][column]Matrix
// Output:
// resVal is the sum matrix
func MatrixAdd[Ord Number](a, b [][]Ord) (resVal [][]Ord, err error) {
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

// MatrixSub adds two 2D matrices of the same size
// Input:
// a, b are two matrices of the form [rows][column]Matrix
// Output:
// resVal is the substraction  matrix
func MatrixSub[Ord Number](a, b [][]Ord) (resVal [][]Ord, err error) {
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
			resVal[i][j] = a[i][j] - b[i][j]
		}
	}
	return resVal, nil
}

// MatrixScalMult implements the scalar multiplication of a matrix
// Input:
// matr is a matrix of the form [rows][column]Matrix
// scal is a scalar value
// Output:
// resVal is the scalar multiblication scal * [][]matr
func MatrixScalMult[Num Number](matr [][]Num, scal Num) (resVal [][]Num) {
	size := MatrixSize(matr)
	rows := size[0]
	columns := size[1]
	// Expand output slice
	for i := 0; i < rows; i++ {
		resVal = append(resVal, make([]Num, columns))
	}
	// multiplication by 0 equals zero matrix
	if scal == 0 {
		return resVal
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			resVal[i][j] = scal * matr[i][j]
		}
	}
	return resVal
}

// IsSquare checks if a matrix is squared (i.e. rows == columns)
func IsSqare[Num Number](input [][]Num) (is bool, matSize [2]int) {
	matSize = MatrixSize(input)
	return matSize[0] == matSize[1], matSize
}

// MatrixSize estimates the size of a 2D matrix
// Input:
// input is the numerical input matrix of the form [rows][column]Matrix
// Output:
// matrixSize is a vector with the size of the matrix [rows, columns]
func MatrixSize[Ord Number](input [][]Ord) (matrixSize [2]int) {
	matrixSize = [2]int{len(input[:][:]), len(input[0][:])}
	return
}

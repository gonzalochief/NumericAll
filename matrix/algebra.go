package matrix

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var ErrDivByZero = errors.New("divide by zero")
var ErrMatSizeMissmatch = errors.New("matrix size missmatch")
var ErrVecSizeMissmatch = errors.New("vector size missmatch")
var ErrMatNotCompatible = errors.New("matrices are not compatible")
var ErrMatNotSquare = errors.New("vector size missmatch")

// Set of all numbers R(real) U C(complex) U i(imaginary)
type Number interface {
	constraints.Float | constraints.Integer | constraints.Complex
}

// Set of Real Numbers
type Real interface {
	constraints.Float | constraints.Integer
}

// MatrixAdd adds two 2D matrices of the same size
// Input:
// a, b are two matrices of the form [rows][column]Matrix
// Output:
// resVal is the sum matrix
func MatrixAdd[Num Number](a, b [][]Num) (resVal [][]Num, err error) {
	sizeA := MatrixSize(a)
	sizeB := MatrixSize(b)
	if sizeA != sizeB {
		return nil, ErrMatSizeMissmatch
	}
	rows := sizeA[0]
	columns := sizeA[1]
	// Expand output slice
	for i := 0; i < rows; i++ {
		resVal = append(resVal, make([]Num, columns))
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
func MatrixSub[Num Number](a, b [][]Num) (resVal [][]Num, err error) {
	sizeA := MatrixSize(a)
	sizeB := MatrixSize(b)
	if sizeA != sizeB {
		return nil, ErrMatSizeMissmatch
	}
	rows := sizeA[0]
	columns := sizeA[1]
	// Expand output slice
	for i := 0; i < rows; i++ {
		resVal = append(resVal, make([]Num, columns))
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

// MatrixMult implements the multiplication of two matrix using the naive approach
// Input:
// a, and b are two compatible matrices of the form [rows][column]Matrix
// Output:
// resVal is the scalar multiblication scal * [][]matr
func MatrixMult[Num Number](a, b [][]Num) (resVal [][]Num, err error) {
	sizeA := MatrixSize(a)
	sizeB := MatrixSize(b)
	if sizeA[1] != sizeB[0] {
		return nil, ErrMatNotCompatible
	} else {
		for i := 0; i < sizeA[0]; i++ {
			resVal = append(resVal, make([]Num, sizeB[1]))
		}
		for i := range a {
			sliceA := GetRow(a, i)
			for j := range b[0] {
				sliceB := GetCol(b, j)
				resVal[i][j], _ = VectScalMult(sliceA, sliceB)
			}
		}
	}
	fmt.Println(resVal)
	return
}

// VectScalMult implements the scalar multiplication of vectors
// Input:
// a, and b are two equaly sized (equal number of components) vectors of the form []Vector
// Output:
// resVal is the scalar multiblication []a * b[]
func VectScalMult[Num Number](a, b []Num) (resVal Num, err error) {
	sizeA := len(a)
	sizeB := len(b)
	if sizeA != sizeB {
		return resVal, ErrVecSizeMissmatch
	}
	for i := 0; i < sizeA; i++ {
		resVal += a[i] * b[i]
	}
	return
}

// GetCol extracts a given column out from a 2D slice (matrix)
func GetCol[Num Number](input [][]Num, col int) (resVal []Num) {
	//check error col number
	for j := range input {
		resVal = append(resVal, input[j][col])
	}
	return
}

// GetRow extracts a given row out from a 2D slice (matrix)
func GetRow[Num Number](input [][]Num, row int) (output []Num) {
	return input[row]
}

// IsSquare checks if a matrix is squared (i.e. rows == columns)
func IsSquare[Num Number](input [][]Num) (is bool, matSize [2]int) {
	matSize = MatrixSize(input)
	return matSize[0] == matSize[1], matSize
}

// MatrixSize estimates the size of a 2D matrix
// Input:
// input is the numerical input matrix of the form [rows][column]Matrix
// Output:
// matrixSize is a vector with the size of the matrix [rows, columns]
func MatrixSize[Num Number](input [][]Num) (matrixSize [2]int) {
	matrixSize = [2]int{len(input[:][:]), len(input[0][:])}
	return
}

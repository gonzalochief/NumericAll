package matrix

import (
	"errors"
	"math"
	"math/cmplx"

	"golang.org/x/exp/constraints"
)

var ErrMatSingular = errors.New("singular matrix")

// MatrixDet returns the Determinant of a square matrix using LU factorization
// Input:
// matrix is a matrix of the form [rows][column]Matrix
// Output:
// det is the determinant of the matrix
func MatrixDetReal[Num Real](matrix [][]Num) (det Num, err error) {
	var index, swapCount int = 1, 0
	var num1, num2, total Num = 1, 1, 1
	det = 1
	// assumes lii=1
	checkSquare, size := IsSquare(matrix)
	tempVect := make([]Num, size[1])
	if !checkSquare {
		return det, ErrMatNotSquare
	}
	for i := 0; i < size[0]; i++ {
		index = i

		// Finding index with a non-zero value
		for index < size[0] && matrix[index][i] == 0 {
			index++
		}
		// if there is not a non-zero value, the determinant is equal to zero
		if index == size[0] {
			return 0, ErrMatSingular
		} else if index != i {
			// Swap rows to move non-zero diagonal element to its propper location in the matrix
			matrix[i], matrix[index] = matrix[index], matrix[i]
			swapCount++
		}
		tempVect = matrix[i]
		for j := i + 1; j < size[0]; j++ {
			num1 = tempVect[i]  // value of the diagonal element
			num2 = matrix[j][i] // value of the nex row's element (in the same column)
			for k := 0; k < size[0]; k++ {
				matrix[j][k] = (num1 * matrix[j][k]) - (num2 * tempVect[k])
			}
			total = total * num1
		}
	}
	det = Num(math.Pow(float64(-1), float64(swapCount))) * det

	for i := 0; i < size[0]; i++ {
		det = det * matrix[i][i]
	}
	return (det / total), nil
}

// MatrixDet returns the Determinant of a square matrix using LU factorization
// Input:
// matrix is a matrix of the form [rows][column]Matrix
// Output:
// det is the determinant of the matrix
func MatrixDetComp[Num constraints.Complex](matrix [][]Num) (det Num, err error) {
	var index int = 1
	var num1, num2, total Num = 1, 1, 1
	swapCount := complex(0, 0)
	det = 1
	// assumes lii=1
	checkSquare, size := IsSquare(matrix)
	tempVect := make([]Num, size[1])
	if !checkSquare {
		return det, ErrMatNotSquare
	}
	for i := 0; i < size[0]; i++ {
		index = i

		// Finding index with a non-zero value
		for index < size[0] && matrix[index][i] == 0 {
			index++
		}
		// if there is not a non-zero value, the determinant is equal to zero
		if index == size[0] {
			return 0, ErrMatSingular
		} else if index != i {
			// Swap rows to move non-zero diagonal element to its propper location in the matrix
			matrix[i], matrix[index] = matrix[index], matrix[i]
			swapCount++
		}
		tempVect = matrix[i]
		for j := i + 1; j < size[0]; j++ {
			num1 = tempVect[i]  // value of the diagonal element
			num2 = matrix[j][i] // value of the nex row's element (in the same column)
			for k := 0; k < size[0]; k++ {
				matrix[j][k] = (num1 * matrix[j][k]) - (num2 * tempVect[k])
			}
			total = total * num1
		}
	}

	det = Num(cmplx.Pow(-1, swapCount)) * det

	for i := 0; i < size[0]; i++ {
		det = det * matrix[i][i]
	}
	return (det / total), nil
}

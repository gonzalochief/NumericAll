package matrix

import (
	"errors"
	"reflect"
	"testing"
)

type testMatrixSum struct {
	TestMatrixAInt [][]int
	TestMatrixBInt [][]int
	TestResMatInt  [][]int
	TestMatrixAF64 [][]float64
	TestMatrixBF64 [][]float64
	TestResMatF64  [][]float64
	ExpectedSize   [2]int
	ExpectedError  error
}

func TestMatrixSum(t *testing.T) {
	testCases := make([]testMatrixSum, 2)
	// Test case - size missmatch error return
	testCases[0].TestMatrixAInt = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	testCases[0].TestMatrixBInt = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	testCases[0].TestMatrixAF64 = [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	testCases[0].TestMatrixBF64 = [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	testCases[0].ExpectedError = ErrMatSizeMissmatch

	// Test case - success
	testCases[1].TestMatrixAInt = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	testCases[1].TestMatrixBInt = [][]int{
		{10, 11, 12},
		{13, 14, 15},
		{16, 17, 18},
	}
	testCases[1].TestResMatInt = [][]int{
		{11, 13, 15},
		{17, 19, 21},
		{23, 25, 27},
	}
	testCases[1].TestMatrixAF64 = [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	testCases[1].TestMatrixBF64 = [][]float64{
		{10, 11, 12},
		{13, 14, 15},
		{16, 17, 18},
	}
	testCases[1].TestResMatF64 = [][]float64{
		{11, 13, 15},
		{17, 19, 21},
		{23, 25, 27},
	}
	testCases[1].ExpectedError = nil

	for _, tc := range testCases {
		resInt, err := MatrixAdd(tc.TestMatrixAInt, tc.TestMatrixBInt)
		if !errors.Is(err, tc.ExpectedError) {
			t.Errorf("failed to detect matrix size missmatch, int variable type")
		}
		if !reflect.DeepEqual(tc.TestResMatInt, resInt) {
			t.Errorf("wrong result value, int variable type")
		}
		resF64, err := MatrixAdd(tc.TestMatrixAF64, tc.TestMatrixBF64)
		if !errors.Is(err, tc.ExpectedError) {
			t.Errorf("failed to detect matrix size missmatch, float variable type")
		}
		if !reflect.DeepEqual(tc.TestResMatF64, resF64) {
			t.Errorf("wrong result value, float variable type")
		}
	}
}

type testStrMatrixSize struct {
	TestMatrixInt [][]int
	TestMatrixF64 [][]float64
	ExpectedSize  [2]int
}

func TestMatrixSize(t *testing.T) {
	testCasesInt := make([]testStrMatrixSize, 2)
	testCasesInt[0].TestMatrixInt = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	testCasesInt[0].TestMatrixF64 = [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	testCasesInt[0].ExpectedSize = [2]int{3, 3}
	testCasesInt[1].TestMatrixInt = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	testCasesInt[1].TestMatrixF64 = [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	testCasesInt[1].ExpectedSize = [2]int{2, 3}
	for _, tc := range testCasesInt {
		size := MatrixSize(tc.TestMatrixInt)
		if size != tc.ExpectedSize {
			t.Errorf("wrong matrix size estimation for int type. Expected: %v, received: %v", tc.ExpectedSize, size)
		}
		size = MatrixSize(tc.TestMatrixF64)
		if size != tc.ExpectedSize {
			t.Errorf("wrong matrix size estimation for float64 type. Expected: %v, received: %v", tc.ExpectedSize, size)
		}
	}
}

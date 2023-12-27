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

func TestMatrixSub(t *testing.T) {
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
		{20, 8, 0},
		{12, 7, 18},
		{10, 16, 17},
	}
	testCases[1].TestMatrixBInt = [][]int{
		{3, 11, 8},
		{2, 17, 12},
		{18, 2, 15},
	}
	testCases[1].TestResMatInt = [][]int{
		{17, -3, -8},
		{10, -10, 6},
		{-8, 14, 2},
	}
	testCases[1].TestMatrixAF64 = [][]float64{
		{20, 8, 0},
		{12, 7, 18},
		{10, 16, 17},
	}
	testCases[1].TestMatrixBF64 = [][]float64{
		{3, 11, 8},
		{2, 17, 12},
		{18, 2, 15},
	}
	testCases[1].TestResMatF64 = [][]float64{
		{17, -3, -8},
		{10, -10, 6},
		{-8, 14, 2},
	}
	testCases[1].ExpectedError = nil

	for _, tc := range testCases {
		resInt, err := MatrixSub(tc.TestMatrixAInt, tc.TestMatrixBInt)
		if !errors.Is(err, tc.ExpectedError) {
			t.Errorf("failed to detect matrix size missmatch, int variable type")
		}
		if !reflect.DeepEqual(tc.TestResMatInt, resInt) {
			t.Errorf("wrong result value, int variable type")
		}
		resF64, err := MatrixSub(tc.TestMatrixAF64, tc.TestMatrixBF64)
		if !errors.Is(err, tc.ExpectedError) {
			t.Errorf("failed to detect matrix size missmatch, float variable type")
		}
		if !reflect.DeepEqual(tc.TestResMatF64, resF64) {
			t.Errorf("wrong result value, float variable type")
		}
	}
}

type testStrMatrixScalMult struct {
	TestScalInt    int
	TestMatrixInt  [][]int
	TestResMatInt  [][]int
	TestScalF64    float64
	TestMatrixF64  [][]float64
	TestResMatF64  [][]float64
	TestScalC128   complex128
	TestMatrixC128 [][]complex128
	TestResMatC128 [][]complex128
	ExpectedError  error
}

func TestMatrixScalMult(t *testing.T) {
	testCases := make([]testStrMatrixScalMult, 2)
	// Test case - size missmatch error return
	testCases[0].TestScalInt = 5
	testCases[0].TestMatrixInt = [][]int{
		{0, 3, 0},
		{0, 1, 3},
		{3, 5, 4},
	}
	testCases[0].TestResMatInt = [][]int{
		{0, 15, 0},
		{0, 5, 15},
		{15, 25, 20},
	}
	testCases[0].TestScalF64 = 5
	testCases[0].TestMatrixF64 = [][]float64{
		{0, 3, 0},
		{0, 1, 3},
		{3, 5, 4},
	}
	testCases[0].TestResMatF64 = [][]float64{
		{0, 15, 0},
		{0, 5, 15},
		{15, 25, 20},
	}
	testCases[0].TestScalC128 = 1 + 3i
	testCases[0].TestMatrixC128 = [][]complex128{
		{4 + 1i, 4 + 4i, 1 + 1i},
		{4 + 1i, 1 + 3i, 3 + 4i},
		{2 + 2i, 2 + 3i, 4 + 3i},
	}
	testCases[0].TestResMatC128 = [][]complex128{
		{1 + 13i, -8 + 16i, -2 + 4i},
		{1 + 13i, -8 + 6i, -9 + 13i},
		{-4 + 8i, -7 + 9i, -5 + 15i},
	}
	testCases[0].ExpectedError = nil
	// Test case: multiplication by 0
	testCases[1].TestScalInt = 0
	testCases[1].TestMatrixInt = [][]int{
		{0, 3, 0},
		{0, 1, 3},
		{3, 5, 4},
	}
	testCases[1].TestResMatInt = [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	testCases[1].TestScalF64 = 0
	testCases[1].TestMatrixF64 = [][]float64{
		{0, 3, 0},
		{0, 1, 3},
		{3, 5, 4},
	}
	testCases[1].TestResMatF64 = [][]float64{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	testCases[1].TestScalC128 = 0
	testCases[1].TestMatrixC128 = [][]complex128{
		{4 + 1i, 4 + 4i, 1 + 1i},
		{4 + 1i, 1 + 3i, 3 + 4i},
		{2 + 2i, 2 + 3i, 4 + 3i},
	}
	testCases[1].TestResMatC128 = [][]complex128{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	testCases[1].ExpectedError = nil

	for _, tc := range testCases {
		resInt := MatrixScalMult(tc.TestMatrixInt, tc.TestScalInt)
		if !reflect.DeepEqual(tc.TestResMatInt, resInt) {
			t.Errorf("wrong result value, int variable type")
		}
		resF64 := MatrixScalMult(tc.TestMatrixF64, tc.TestScalF64)
		if !reflect.DeepEqual(tc.TestResMatF64, resF64) {
			t.Errorf("wrong result value, float variable type")
		}
		resC128 := MatrixScalMult(tc.TestMatrixC128, tc.TestScalC128)
		if !reflect.DeepEqual(tc.TestResMatC128, resC128) {
			t.Errorf("wrong result value, complex 128 variable type")
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

package matrix

import (
	"errors"
	"testing"
)

type testStrsDet struct {
	TestMatrF64   [][]float64
	ExpDetF64     float64
	TestMatrInt   [][]int
	ExpDetInt     int
	ExpectedError error
}

func TestDetFunc(t *testing.T) {
	testCase := make([]testStrsDet, 3)
	// Test cases: success
	testCase[0].TestMatrF64 = [][]float64{
		{1, 0, 2, -1},
		{3, 0, 0, 5},
		{2, 1, 4, -3},
		{1, 0, 5, 0},
	}
	testCase[0].ExpDetF64 = 30
	testCase[0].TestMatrInt = [][]int{
		{1, 0, 2, -1},
		{3, 0, 0, 5},
		{2, 1, 4, -3},
		{1, 0, 5, 0},
	}
	testCase[0].ExpDetInt = 30
	testCase[1].TestMatrF64 = [][]float64{
		{2, -1, 3, 0},
		{4, -2, 7, 0},
		{-3, -4, 1, 5},
		{6, -6, 8, 0},
	}
	testCase[1].ExpDetF64 = -30
	testCase[1].TestMatrInt = [][]int{
		{2, -1, 3, 0},
		{4, -2, 7, 0},
		{-3, -4, 1, 5},
		{6, -6, 8, 0},
	}
	testCase[1].ExpDetInt = -30
	testCase[2].TestMatrF64 = [][]float64{
		{2, 1, -1, 1},
		{1, 1, 0, 3},
		{-1, 2, 3, -1},
		{3, -1, -1, 2},
	}
	testCase[2].ExpDetF64 = 39
	testCase[2].TestMatrInt = [][]int{
		{2, 1, -1, 1},
		{1, 1, 0, 3},
		{-1, 2, 3, -1},
		{3, -1, -1, 2},
	}
	testCase[2].ExpDetInt = 39
	// Test cases: fail - nota a square matrix
	// Test cases: fail - nota an invertible matrix (singular matrix)
	for _, tc := range testCase {
		detF64, err := MatrixDetReal(tc.TestMatrF64)
		t.Log(err)
		t.Log(detF64)
		if err == nil {
			if detF64 != tc.ExpDetF64 {
				t.Errorf("wrong result value, complex128 variable type")
			}
		} else if !errors.Is(err, tc.ExpectedError) {
			t.Errorf("failed to detect error, complex128 variable type")
		}
		detInt, err := MatrixDetReal(tc.TestMatrInt)
		t.Log(err)
		t.Log(detInt)
		if err == nil {
			if float64(detInt) != float64(tc.ExpDetInt) {
				t.Errorf("wrong result value, complex128 variable type")
			}
		} else if !errors.Is(err, tc.ExpectedError) {
			t.Errorf("failed to detect error, complex128 variable type")
		}
	}

}


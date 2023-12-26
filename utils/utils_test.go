package utils

import (
	"math"
	"testing"
)

// TestInverseSqrt tests the implementation of the fast inverse square root function (Quake III) for float (32 and 64 bits) (with generics)
// Assumptions:
// - Tests estimation against a maximum error of 5 decimal points (epsilon of 1e-5)
func TestInverseSqrt(t *testing.T) {
	testCases32 := []struct {
		input         float32
		expected      float32
		iterations    int
		expectedError error
	}{
		{1, 1, 2, nil}, {4, 0.5, 2, nil}, {16, 0.25, 2, nil}, {25, 0.2, 2, nil}, {64, 0.125, 2, nil}, {100, 0.1, 2, nil}, {256, 0.0625, 2, nil}, {400, 0.05, 2, nil}, {625, 0.04, 2, nil}, {1024, 0.03125, 2, nil}, {1024, 0.03125, 0, ErrIterMin}, {-1024, NaN[float32](), 2, ErrIndRes},
	}

	epsilon32 := float32(1e-5)

	for _, tc := range testCases32 {
		testValue, err := FastInvSqrt(tc.input, tc.iterations)
		if err != tc.expectedError {
			t.Error("iteration is not equal or greater than 1")
			break
		} else if Abs(testValue-tc.expected) > epsilon32 {
			t.Errorf("invalid result. Expecting: %.10f, Receiving: %.10f, Difference: %.10f", tc.expected, testValue, -tc.expected+testValue)
		}
	}

	testCases64 := []struct {
		input         float64
		expected      float64
		iterations    int
		expectedError error
	}{
		{1, 1, 2, nil}, {4, 0.5, 2, nil}, {16, 0.25, 2, nil}, {25, 0.2, 2, nil}, {64, 0.125, 2, nil}, {100, 0.1, 2, nil}, {256, 0.0625, 2, nil}, {400, 0.05, 2, nil}, {625, 0.04, 2, nil}, {1024, 0.03125, 2, nil}, {1024, 0.03125, 0, ErrIterMin}, {-1024, NaN[float64](), 2, ErrIndRes},
	}

	epsilon64 := float64(1e-5)

	for _, tc := range testCases64 {
		testValue, err := FastInvSqrt(tc.input, tc.iterations)
		if err != tc.expectedError {
			t.Error("iteration is not equal or greater than 1")
			break
		} else if Abs(testValue-tc.expected) > epsilon64 {
			t.Errorf("invalid result. Expecting: %.10f, Receiving: %.10f, Difference: %.10f", tc.expected, testValue, -tc.expected+testValue)
		}
	}
}

// TestInverseSqrt32 tests the implementation of the fast inverse square root function (Quake III) for float32
// Assumptions:
// - Tests estimation against a maximum error of 5 decimal points (epsilon of 1e-5)
func TestInverseSqrt32(t *testing.T) {
	testCases := []struct {
		input         float32
		expected      float32
		iterations    int
		expectedError error
	}{
		{1, 1, 2, nil}, {4, 0.5, 2, nil}, {16, 0.25, 2, nil}, {25, 0.2, 2, nil}, {64, 0.125, 2, nil}, {100, 0.1, 2, nil}, {256, 0.0625, 2, nil}, {400, 0.05, 2, nil}, {625, 0.04, 2, nil}, {1024, 0.03125, 2, nil}, {1024, 0.03125, 0, ErrIterMin}, {-1024, NaN[float32](), 2, ErrIndRes},
	}

	epsilon := float32(1e-5)

	for _, tc := range testCases {
		testValue, err := FastInvSqrt32(tc.input, tc.iterations)
		if err != tc.expectedError {
			t.Error("iteration is not equal or greater than 1")
			break
		} else if Abs(testValue-tc.expected) > epsilon {
			t.Errorf("invalid result. Expecting: %.10f, Receiving: %.10f, Difference: %.10f", tc.expected, testValue, -tc.expected+testValue)
		}
	}

}

// TestInverseSqrt64 tests the implementation of the fast inverse square root function (Quake III) for float32
// Assumptions:
// - Tests estimation against a maximum error of 5 decimal points (epsilon of 1e-5)
func TestInverseSqrt64(t *testing.T) {
	testCases := []struct {
		input         float64
		expected      float64
		iterations    int
		expectedError error
	}{
		{1, 1, 2, nil}, {4, 0.5, 2, nil}, {16, 0.25, 2, nil}, {25, 0.2, 2, nil}, {64, 0.125, 2, nil}, {100, 0.1, 2, nil}, {256, 0.0625, 2, nil}, {400, 0.05, 2, nil}, {625, 0.04, 2, nil}, {1024, 0.03125, 2, nil}, {1024, 0.03125, 0, ErrIterMin}, {-1024, NaN[float64](), 2, ErrIndRes},
	}

	epsilon := float64(1e-5)

	for _, tc := range testCases {
		testValue, err := FastInvSqrt64(tc.input, tc.iterations)
		if err != tc.expectedError {
			t.Error("iteration is not equal or greater than 1")
			break
		} else if Abs(testValue-tc.expected) > epsilon {
			t.Errorf("invalid result. Expecting: %.10f, Receiving: %.10f, Difference: %.10f", tc.expected, testValue, -tc.expected+testValue)
		}
	}

}

type testStrAbs struct {
	TestValF64 []float64
	TestResF64 []float64
	TestValF32 []float32
	TestResF32 []float32
}

func TestAbsFloat(t *testing.T) {
	// Set test cases
	testCasesAbs := make([]testStrAbs, 1)
	testCasesAbs[0].TestValF64 = []float64{-1000, -500, -0.1, 0, 0.1, 500, 1000, math.Inf(1), math.Inf(-1)}
	testCasesAbs[0].TestResF64 = []float64{1000, 500, 0.1, 0, 0.1, 500, 1000, math.Inf(1), math.Inf(1)}
	testCasesAbs[0].TestValF32 = []float32{-1000, -500, -0.1, 0, 0.1, 500, 1000, float32(math.Inf(1)), float32(math.Inf(-1))}
	testCasesAbs[0].TestResF32 = []float32{1000, 500, 0.1, 0, 0.1, 500, 1000, float32(math.Inf(1)), float32(math.Inf(1))}
	// Test case: algorythm prformance and results
	for _, tc := range testCasesAbs {
		for i, tVal := range tc.TestValF64 {
			if Abs(tVal) != tc.TestResF64[i] {
				t.Errorf("wrong value for float64, expected: %f, received: %f", tc.TestResF64[i], Abs(tVal))
			}
		}
		for i, tVal := range tc.TestValF32 {
			if Abs(tVal) != tc.TestResF32[i] {
				t.Errorf("wrong value for float32, expected: %f, received: %f", tc.TestResF32[i], Abs(tVal))
			}
		}
	}
	// Test case: NaN special case
	if !IsNaN(Abs(NaN[float64]())) {
		t.Errorf("wrong value for float64, expected: NaN, received: %f", math.NaN())
	}
	if !IsNaN(Abs(NaN[float32]())) {
		t.Errorf("wrong value for float32, expected: NaN, received: %f", float32(math.NaN()))
	}
}

func BenchmarkAbsFloat(b *testing.B) {
	// Set test cases
	testCasesAbs := make([]testStrAbs, 1)
	testCasesAbs[0].TestValF64 = []float64{-1000, -500, -0.1, 0, 0.1, 500, 1000, math.Inf(1), math.Inf(-1)}
	testCasesAbs[0].TestResF64 = []float64{1000, 500, 0.1, 0, 0.1, 500, 1000, math.Inf(1), math.Inf(1)}
	testCasesAbs[0].TestValF32 = []float32{-1000, -500, -0.1, 0, 0.1, 500, 1000, float32(math.Inf(1)), float32(math.Inf(-1))}
	testCasesAbs[0].TestResF32 = []float32{1000, 500, 0.1, 0, 0.1, 500, 1000, float32(math.Inf(1)), float32(math.Inf(1))}
	// Test case: algorythm prformance and results
	for _, tc := range testCasesAbs {
		for _, tVal := range tc.TestValF64 {
			_ = Abs(tVal)
		}
	}
}

func BenchmarkAbsFloatGo(b *testing.B) {
	// Set test cases
	testCasesAbs := make([]testStrAbs, 1)
	testCasesAbs[0].TestValF64 = []float64{-1000, -500, -0.1, 0, 0.1, 500, 1000, math.Inf(1), math.Inf(-1)}
	testCasesAbs[0].TestResF64 = []float64{1000, 500, 0.1, 0, 0.1, 500, 1000, math.Inf(1), math.Inf(1)}
	testCasesAbs[0].TestValF32 = []float32{-1000, -500, -0.1, 0, 0.1, 500, 1000, float32(math.Inf(1)), float32(math.Inf(-1))}
	testCasesAbs[0].TestResF32 = []float32{1000, 500, 0.1, 0, 0.1, 500, 1000, float32(math.Inf(1)), float32(math.Inf(1))}
	// Test case: algorythm prformance and results
	for _, tc := range testCasesAbs {
		for _, tVal := range tc.TestValF64 {
			_ = math.Abs(tVal)
		}
	}
}

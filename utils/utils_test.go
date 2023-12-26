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

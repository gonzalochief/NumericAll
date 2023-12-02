package utils

import (
	"math"
	"testing"
)

// TestInverseSqrt tests the implementation of the fast inverse square root function (Quake III) for float32
// Assumptions:
// - Tests estimation against a maximum error of 5 decimal points (epsilon of 1e-5)
func TestInverseSqrt32(t *testing.T) {
	testCases := []struct {
		input    float32
		expected float32
	}{
		{1, 1}, {4, 0.5}, {16, 0.25}, {25, 0.2}, {64, 0.125}, {100, 0.1}, {256, 0.0625}, {400, 0.05}, {625, 0.04}, {1024, 0.03125},
	}

	epsilon := float64(1e-5)

	for _, tc := range testCases {
		testValue, err := FastInvSqrt32(tc.input, 2)
		if err != nil {
			t.Error("iteration is not equal or greater than 1")
			break
		} else if math.Abs(float64(testValue)-float64(tc.expected)) > epsilon {
			t.Errorf("invalid result. Expecting: %.10f, Receiving: %.10f, Difference: %.10f", tc.expected, testValue, -tc.expected+testValue)
		}
	}

}

// TestInverseSqrt tests the implementation of the fast inverse square root function (Quake III) for float32
// Assumptions:
// - Tests estimation against a maximum error of 5 decimal points (epsilon of 1e-5)
func TestInverseSqrt64(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{1, 1}, {4, 0.5}, {16, 0.25}, {25, 0.2}, {64, 0.125}, {100, 0.1}, {256, 0.0625}, {400, 0.05}, {625, 0.04}, {1024, 0.03125},
	}

	epsilon := float64(1e-5)

	for _, tc := range testCases {
		testValue, err := FastInvSqrt64(tc.input, 2)
		if err != nil {
			t.Error("iteration is not equal or greater than 1")
			break
		} else if math.Abs(float64(testValue)-float64(tc.expected)) > epsilon {
			t.Errorf("invalid result. Expecting: %.10f, Receiving: %.10f, Difference: %.10f", tc.expected, testValue, -tc.expected+testValue)
		}
	}

}

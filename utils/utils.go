package utils

import (
	"errors"
	"math"

// FastInvSqrt implements Quake III's Fast Inverse Square Root algorythm for Float types (64 and 32)
// Input:
// - input: float radicand
// - iterations: integer with the required estimation iterations (Newton-Raphson approximation)
// Output:
// - output: inverse square root
// - error
func FastInvSqrt[T constraints.Float](input T, iterations int) (output T, err error) {
	if input < 0 {
		return NaN[T](), ErrIndRes
	}
	if iterations > 0 {
		var x2 T = input * 0.5
		switch any(input).(type) {
		case float64:
			i := math.Float64bits(float64(input))
			i = 0x5FE6EB50C7B537A9 - (i >> 1) // Robertson estimation of the Magic Number
			output = T(math.Float64frombits(i))
			output = output * (1.5 - (x2 * output * output))
			for j := 1; j < iterations; j++ {
				output = output * (1.5 - (x2 * output * output))
			}
		case float32:
			i := math.Float32bits(float32(input))
			//i = 0x5F3759DF - (i >> 1) // Original Quake III estimation  of the Magic Number
			i = 0x5F375A86 - (i >> 1) // Lomont estimation  of the Magic Number
			output = T(math.Float32frombits(i))
			output = output * (1.5 - (x2 * output * output))
			for j := 1; j < iterations; j++ {
				output = output * (1.5 - (x2 * output * output))
			}
		}
	} else {
		return NaN[T](), ErrIterMin
	}
	return output, nil
}

// FastInvSqrt32 implements Quake III's Fast Inverse Square Root algorythm for float32 types
// Input:
// - input: float 32 radicand
// - iterations: integer with the required estimation iterations (Newton-Raphson approximation)
// Output:
// - output: inverse square root
// - error
func FastInvSqrt32(input float32, iterations int) (output float32, err error) {
	if input < 0 {
		return float32(math.NaN()), errors.New("indetermined square root. negative input")
	}
	if iterations > 0 {
		var x2 float32
		x2 = input * 0.5
		i := math.Float32bits(input)
		//i = 0x5F3759DF - (i >> 1) // Original Quake III estimation  of the Magic Number
		i = 0x5F375A86 - (i >> 1) // Lomont estimation  of the Magic Number
		output = math.Float32frombits(i)
		output = output * (1.5 - (x2 * output * output))
		for i := 1; i < iterations; i++ {
			output = output * (1.5 - (x2 * output * output))
			break
		}
	} else {
		return float32(math.NaN()), errors.New("iterations must be equal or greater that 1")
	}
	return output, nil
}

// FastInvSqrt64 implements Quake III's Fast Inverse Square Root algorythm for float32 types
// Input:
// - input: float 64 radicand
// - iterations: integer with the required estimation iterations (Newton-Raphson approximation)
// Output:
// - output: inverse square root
// - error
func FastInvSqrt64(input float64, iterations int) (output float64, err error) {
	if input < 0 {
		return math.NaN(), errors.New("indetermined square root. negative input")
	}
	if iterations > 0 {
		var x2 float64
		x2 = input * 0.5
		i := math.Float64bits(input)
		i = 0x5FE6EB50C7B537A9 - (i >> 1) // Robertson estimation of the Magic Number
		output = math.Float64frombits(i)
		output = output * (1.5 - (x2 * output * output))
		for i := 1; i < iterations; i++ {
			output = output * (1.5 - (x2 * output * output))
			break
		}
	} else {
		return math.NaN(), errors.New("iterations must be equal or greater that 1")
	}
	return output, nil
}

// Abs returns the absolute value of x.
//
// Special cases are:
//
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs[T constraints.Float](x T) T {
	var getTemp T
	switch any(x).(type) {
	case float64:
		getTemp = T(math.Float64frombits(math.Float64bits(any(x).(float64)) &^ (1 << 63)))
	case float32:
		getTemp = T(math.Float32frombits(math.Float32bits(any(x).(float32)) &^ (1 << 31)))
	}
	return getTemp
}

// IsNaN reports whether f is an IEEE 754 “not-a-number” value.
func IsNaN[T constraints.Float](f T) (is bool) {
	// IEEE 754 says that only NaNs satisfy f != f.
	// To avoid the floating-point hardware, could use:
	//	x := Float64bits(f);
	//	return uint32(x>>shift)&mask == mask && x != uvinf && x != uvneginf
	return f != f
}

const (
	uvnanDouble = 0x7FF8000000000001
	uvnanSingle = 0x7F800001
)

// NaN returns an IEEE 754 “not-a-number” value.
func NaN[T constraints.Float]() (out T) {
	switch any(out).(type) {
	case float64:
		out = T(math.Float64frombits(uvnanDouble))
	case float32:
		out = T(math.Float32frombits(uvnanSingle))
	}
	return
}

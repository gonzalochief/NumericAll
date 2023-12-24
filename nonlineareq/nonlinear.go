package nonlineareq

import (
	"errors"
	"math"
)

var ErrMaxIter = errors.New("maximum number of iretations reached")
var ErrFuncSignNotEqual = errors.New("the signs of y(a) and y(b) are not different")

// YEqFuncx function type is used to create y=f(x) type of functions
// the params ..float64 variable allows the user to enter configuration parameters to standard funcions (e.g. for irr estimations or any standard function which parameters change case by case).
type YEqFuncx func(x float64) float64

// FixPt estimates the solution to the equation x = g(x) using the p(n+1) = g(p(n)) iteration, which is estimated based on an initial point
// Inputs:
//
//	y is the iteration function
//	p0 is the starting point
//	tol is the tolerance in decimal places
//	maxIter is the maximum number of allowed iterations
//
// Outputs:
//
//	i last iteration
//	pAprox fixed point approximation
//	errAprox Absolute error
//	relErr relative error
//	pSeries fixed point iterations
func FixPt(y YEqFuncx, p0 float64, tol int, maxIter int) (i int, pAprox, errAprox, relErr float64, pSeries []float64, err error) {
	TolDec := float64(1) / math.Pow10(tol)
	epsilon := (math.Nextafter(1, 2) - 1)
	pSeries = append(pSeries, p0)
	for i = 1; i <= maxIter; i++ {
		pSeries = append(pSeries, y(pSeries[i-1]))
		errAprox = math.Abs(pSeries[i] - pSeries[i-1])
		relErr = errAprox / (math.Abs(pSeries[i]) + epsilon)
		pAprox = pSeries[i]
		if (errAprox < TolDec) || (relErr < TolDec) {
			return i, pAprox, errAprox, relErr, pSeries, nil
		}
	}
	return 0, 0, 0, 0, nil, ErrMaxIter
}

// BisectBolzano estimates the value of x that makes the function equal to 0 inside the interval [a,b] using the Bolzano's Bisection Method
// The method only works if the values of f(a) and f(b) have different signs
// Inputs:
//
//		y is a function of the form:
//	    f(x) = expression
//	 that has been reexpressed as 0 = expression - f(x)
//
//	 a and b are the left and right extreme values of the interval
//		tol is the tolerance
//
// Outputs:
//
//	c is the zero
//	yC is the function value evaluated at c
//	absErr is the error of the approximation
func BisectBolzano(y YEqFuncx, a, b, tol float64) (c, yC, absErr float64, err error) {
	var i int
	ya := y(a)
	yb := y(b)
	if ya*yb > 0 {
		return 0, 0, 0, ErrFuncSignNotEqual
	}
	maxIter := 1 + math.Round((math.Log(b-a)-math.Log(tol))/math.Log(2))
	for i = 0; i < int(maxIter); i++ {
		c = (a + b) / 2
		yC = y(c)
		if yC == 0 {
			a = c
			b = c
		} else if (yb * yC) > 0 {
			b = c
			yb = yC
		} else {
			a = c
			// ya = yC // It does not affect the calculation. Used only for debugging
		}
		if (b - a) < tol {
			c = (a + b) / 2
			return c, y(c), math.Abs(b - a), nil
		}
	}
	return 0, 0, 0, ErrMaxIter
}

// RegulaFalsi estimates the value of x that makes the function equal to 0 inside the interval [a,b] using the Régula Falsi Method
// The method only works if the values of f(a) and f(b) have different signs
// Inputs:
//
//	y is the function function
//	a and b are the left and right extreme values of the interval
//	tol is the tolerance for the zero
//	epsilon is the tolerance for f(c)
//
// Outputs:
//
//	c is the zero
//	yC is the function value evaluated at c
//	absErr is the error of the approximation
func RegulaFalsi(y YEqFuncx, a, b, tol, epsilon float64, maxIter int) (c, yC, absErr float64, err error) {
	var dx, ac float64
	ya := y(a)
	yb := y(b)
	if ya*yb > 0 {
		return 0, 0, 0, ErrFuncSignNotEqual
	}
	for i := 0; i < maxIter; i++ {
		dx = yb * (b - a) / (yb - ya)
		c = b - dx
		ac = c - a
		yC = y(c)
		if yC == 0 {
			return c, yC, absErr, nil
		} else if yb*yC > 0 {
			b = c
			yb = yC
		} else {
			a = c
			ya = yC
		}
		dx = math.Min(math.Abs(dx), ac)
		if (math.Abs(dx) < tol) || (math.Abs(yC) < epsilon) {
			return c, yC, absErr, nil
		}
	}
	return 0, 0, 0, ErrMaxIter
}

// NewtonRaphson estimates the value of x that makes the function equal to 0 using the Newton-Raphson Method
// Inputs:
//
//		y is the function function y=f(x)
//		dy is the dirivative of the function function y
//		p0 is the initial point for the zero approximation
//		delta is the tolerance for the zero
//		epsilon is the tolerance for f(c)
//	 maxIter is the maximum iteration for the algorithm
//
// Outputs:
//
//	zeroApr is the approximation to the zero (P0)
//	yC is the function value evaluated at P0
//	absErr is the error of the approximation
//	i is the iteration that generated the approximation
func NewtonRaphson(y, dy YEqFuncx, p0, delta, epsilon float64, maxIter int) (zeroApr, yZero, absErr float64, i int, err error) {
	for i = 0; i < maxIter; i++ {
		p1 := p0 - y(p0)/dy(p0)
		absErr = math.Abs(p1 - p0)
		relErr := 2 * absErr / (math.Abs(p1) + delta)
		p0 = p1
		yP0 := y(p0)
		if (absErr < delta) || (relErr < delta) || (math.Abs(yP0) < epsilon) {
			return p0, yP0, absErr, i, nil
		}
	}
	return math.NaN(), math.NaN(), math.NaN(), i, ErrMaxIter
}

// Secant estimates the value of x that makes the function equal to 0 using the Secant Method
// Inputs:
//
//		y is the function function y=f(x)
//		dy is the dirivative of the function function y
//		p0 and P1 are the initial point for the zero approximation
//		delta is the tolerance for the zero
//		epsilon is the tolerance for f(c)
//	 maxIter is the maximum iteration for the algorithm
//
// Outputs:
//
//	zeroApr is the approximation to the zero (P0)
//	yC is the function value evaluated at P0
//	absErr is the error of the approximation
//	i is the iteration that generated the approximation
func Secant(y YEqFuncx, p0, p1, delta, epsilon float64, maxIter int) (zeroApr, yZero, absErr float64, i int, err error) {
	for i = 0; i < maxIter; i++ {
		p2 := p1 - y(p1)*(p1-p0)/(y(p1)-y(p0))
		absErr = math.Abs(p2 - p1)
		relErr := 2 * absErr / (math.Abs(p2) + delta)
		p0 = p1
		p1 = p2
		yZero = y(p1)
		if (absErr < delta) || (relErr < delta) || (math.Abs(yZero) < epsilon) {
			return p1, yZero, absErr, i, nil
		}
	}
	return math.NaN(), math.NaN(), math.NaN(), i, ErrMaxIter
}

func Steffensen() {}

func Muller() {}

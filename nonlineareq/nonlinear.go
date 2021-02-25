package nonlineareq

import (
	"errors"
	"math"
)

//YEqFuncx function type is used to create y=f(x) type of functions
//the params ..[]float64 variable allows the user to enter configuration parameters to standard funcions (e.g. for irr estimations or any standard function which parameters change case by case).
//for example: a*x+b*x+c*x² should receive a float array of the form [a, b, c], and the fuction shoud be of the form params[0][0]*x+params[0][1]*x+params[0][3]*x²
type YEqFuncx func(x float64, params ...[]float64) float64

//FixPt estimates the solution to the equation x = g(x) using the p(n+1) = g(p(n)) iteration, which is estimated based on an initial point
//Inputs:
//	y is the iteration function
//	p0 is the starting point
//	tol is the tolerance in decimal places
//	maxIter is the maximum number of allowed iterations
//Outputs:
//	i last iteration
//	pAprox fixed point approximation
//	errAprox Absolute error
//	relErr relative error
//	pSeries fixed point iterations
func FixPt(y YEqFuncx, p0 float64, tol int, maxIter int, params ...[]float64) (i int, pAprox, errAprox, relErr float64, pSeries []float64, err error) {
	var TolDec float64
	TolDec = float64(1) / math.Pow10(tol)
	epsilon := (math.Nextafter(1, 2) - 1)
	pSeries = append(pSeries, p0)
	for i = 1; i <= maxIter; i++ {
		if params != nil {
			pSeries = append(pSeries, y(pSeries[i-1], params[0]))
		} else {
			pSeries = append(pSeries, y(pSeries[i-1]))
		}
		errAprox = math.Abs(pSeries[i] - pSeries[i-1])
		relErr = errAprox / (math.Abs(pSeries[i]) + epsilon)
		pAprox = pSeries[i]
		if (errAprox < TolDec) || (relErr < TolDec) {
			return i, pAprox, errAprox, relErr, pSeries, nil
		}
	}
	err = errors.New("Maximum number of iretations reached")
	return 0, 0, 0, 0, nil, err
}

//BisectBolzano estimates the value of x that makes the function equal to 0 inside the interval [a,b] using the Bolzano's Bisection Method
//The method only works if the values of f(a) and f(b) have different signs
//Inputs:
//	y is the function function
//	a and b are the left and right extreme values of the interval
//	tol is the tolerance
//Outputs:
//	c is the zero
//	yC is the function value evaluated at c
//	absErr is the error of the approximation
func BisectBolzano(y YEqFuncx, a, b, tol float64, params ...[]float64) (c, yC, absErr float64, err error) {
	var i int
	var ya, yb float64
	if params != nil {
		ya = y(a, params[0])
		yb = y(b, params[0])
	} else {
		ya = y(a)
		yb = y(b)
	}
	if ya*yb > 0 {
		err = errors.New("The signs of a and b are not different")
		return 0, 0, 0, err
	}
	maxIter := 1 + math.Round((math.Log(b-a)-math.Log(tol))/math.Log(2))
	for i = 0; i < int(maxIter); i++ {
		c = (a + b) / 2
		if params != nil {
			yC = y(c, params[0])
		} else {
			yC = y(c)
		}
		if yC == 0 {
			a = c
			b = c
		} else if (yb * yC) > 0 {
			b = c
			yb = yC
		} else {
			a = c
			ya = yC
		}
		if (b - a) < tol {
			c = (a + b) / 2
			return c, math.Abs(b - a), y(c), nil
		}
	}
	err = errors.New("Maximum number of iretations reached")
	return 0, 0, 0, err
}

//RegulaFalsi estimates the value of x that makes the function equal to 0 inside the interval [a,b] using the Régula Falsi Method
//The method only works if the values of f(a) and f(b) have different signs
//Inputs:
//	y is the function function
//	a and b are the left and right extreme values of the interval
//	tol is the tolerance for the zero
//	epsilon is the tolerance for f(c)
//Outputs:
//	c is the zero
//	yC is the function value evaluated at c
//	absErr is the error of the approximation
func RegulaFalsi(y YEqFuncx, a, b, tol, epsilon float64, maxIter int, params ...[]float64) (c, yC, absErr float64, err error) {
	var dx, ac float64
	var ya, yb float64
	if params != nil {
		ya = y(a, params[0])
		yb = y(b, params[0])
	} else {
		ya = y(a)
		yb = y(b)
	}
	if ya*yb > 0 {
		err = errors.New("The signs of a and b are not different")
		return 0, 0, 0, err
	}
	for i := 0; i < int(maxIter); i++ {
		dx = yb * (b - a) / (yb - ya)
		c = b - dx
		ac = c - a
		if params != nil {
			yC = y(c, params[0])
		} else {
			yC = y(c)
		}
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
	err = errors.New("Maximum number of iretations reached")
	return 0, 0, 0, err
}

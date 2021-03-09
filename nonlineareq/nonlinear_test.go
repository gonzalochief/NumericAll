package nonlineareq

import (
	"math"
	"testing"
)

func TestFixedPoint(t *testing.T) {
	var funcToEval YEqFuncx = func(x float64, params ...[]float64) float64 {
		return -4 + 4*x - (1.0/2.0)*math.Pow(x, 2)
	}
	_, pAprox, _, _, _, _ := FixPt(funcToEval, 3.8, 10, 50) //this test case converges to 4
	if pAprox != 4 {
		t.Errorf("The result %f is different to 4", pAprox)
	}
	var y YEqFuncx = func(x float64, params ...[]float64) float64 {
		return -params[0][0] + params[0][1]*x - (1.0/2.0)*math.Pow(x, 2)
	}
	test := []float64{4.0, 4.0}
	c, _ := Steffensen(y, 3.8, 0.000000001, 0.000000001, 50, test)
	if c != 4 {
		t.Errorf("The result %f is different to 4", pAprox)
	}

	var funcToEval2 YEqFuncx = func(x float64, params ...[]float64) float64 {
		return 0.5*x + 1.5
	}
	_, _, absErr, relErr, _, _ := FixPt(funcToEval2, 4, 10, 50) //this test case converges to 3, but with decimals (not an exact solution). The ok trigger commes from the errors (relative) being lower than the tolerance
	if (absErr > (1.0 / math.Pow10(10))) && (relErr > (1.0 / math.Pow10(10))) {
		t.Error("None of the error measures fulfill the tollerance requirement")
	}
}

func TestBisect(t *testing.T) {
	var funcToEval YEqFuncx = func(x float64, params ...[]float64) float64 {
		return x*math.Sin(x) - 1
	}
	c, _, _, _ := BisectBolzano(funcToEval, 0, 2, 0.01)
	if c != 1.11328125 {
		t.Errorf("The result %f is different to 1.11328125", c)
	}
}

func TestRegulafalsi(t *testing.T) {
	var funcToEval YEqFuncx = func(x float64, params ...[]float64) float64 {
		return x*math.Sin(x) - 1
	}
	c, _, _, _ := RegulaFalsi(funcToEval, 0, 2, 0.001, 0.001, 50)
	if c != 1.1141611949626338 {
		t.Errorf("The result %f is different to 1.11328125", c)
	}
}

func TestNewtonRaphson(t *testing.T) {
	var irrFunc YEqFuncx = func(x float64, params ...[]float64) float64 {
		lenFcf := len(params[0])
		var npv float64 = 0
		for i := 0; i < lenFcf; i++ {
			npv += params[0][i] / math.Pow(1+x, float64(i))
		}
		return npv
	}

	//
	var dxirrFunc YEqFuncx = func(x float64, params ...[]float64) float64 {
		lenFcf := len(params[0])
		var npv float64 = 0
		for i := 0; i < lenFcf; i++ {
			npv += -(float64(i) * params[0][i]) / math.Pow(1+x, float64(i)+1)
		}
		return npv
	}
	var fcf = []float64{-1000, 10, 10, 10, 10, 10, 10, 2000, 100}
	c, _, _, _, _ := NewtonRaphson(irrFunc, dxirrFunc, 0.0, 0.000001, 0.00000001, 50, fcf)
	if c != 0.11774096121389947 {
		t.Errorf("The result %f is different to 1.11328125", c)
	}
}

func TestNewtonRaphsonAccel(t *testing.T) {
	var y YEqFuncx = func(x float64, params ...[]float64) float64 {
		return math.Pow(x, 3) - 3.0*x + 2
	}

	//
	var dy YEqFuncx = func(x float64, params ...[]float64) float64 {
		return 3.0*math.Pow(x, 2) - 3.0
	}
	_, _, absErr, _, _ := NewtonRaphsonAccel(y, dy, 0.0, 0.00000001, 0.0000000001, 2, 50)
	if absErr > 1.0000416614589853 {
		t.Errorf("The result %f is different to 1.0000416614589853", absErr)
	}
}

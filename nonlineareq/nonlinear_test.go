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

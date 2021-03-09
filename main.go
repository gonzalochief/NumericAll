package main

import (
	"fmt"
	"math"

	"github.com/gonzalochief/NumericAll/nonlineareq"
)

func main() {

	var funcToEval nonlineareq.YEqFuncx = func(x float64, params ...[]float64) float64 {
		return -params[0][0] + params[0][1]*x - (1.0/2.0)*math.Pow(x, 2)
	}
	test := []float64{4.0, 4.0}
	s := funcToEval(10, test)
	fmt.Println("Func: ", s)
	i, pApr, errApr, relErr, pSlice, _ := nonlineareq.FixPt(funcToEval, 3.8, 10, 50, test)
	fmt.Println("i: ", i)
	fmt.Println("p aprox.: ", pApr)
	fmt.Println("error aprox.: ", errApr)
	fmt.Println("relative err: ", relErr)
	fmt.Println("p slice: ", pSlice)

	fmt.Println("Steffensson method")
	c, errsf := nonlineareq.Steffensen(funcToEval, 3.8, 0.000000001, 0.000000001, 50, test)
	fmt.Println("c estimated: ", c)
	fmt.Println("err: ", errsf)

	var funcToEval2 nonlineareq.YEqFuncx = func(x float64, params ...[]float64) float64 {
		return 0.5*x + 1.5
	}
	i, pApr, errApr, relErr, pSlice, _ = nonlineareq.FixPt(funcToEval2, 4, 10, 50)
	fmt.Println("i: ", i)
	fmt.Println("p aprox.: ", pApr)
	fmt.Println("error aprox.: ", errApr)
	fmt.Println("relative err: ", relErr)
	fmt.Println("p slice: ", pSlice)

	fmt.Println("Bisection method")
	var funcToEval3 nonlineareq.YEqFuncx = func(x float64, params ...[]float64) float64 {
		return x*math.Sin(x) - 1
	}
	fmt.Println("evaluada: ", funcToEval3(0))
	c, yC, errApr, _ := nonlineareq.BisectBolzano(funcToEval3, 0, 2, 0.01)
	fmt.Println("c: ", c)
	fmt.Println("yC aprox.: ", yC)
	fmt.Println("error aprox.: ", errApr)

	fmt.Println("Regula Falsi method")
	c, yC, errApr, _ = nonlineareq.RegulaFalsi(funcToEval3, 0, 2, 0.001, 0.001, 50)
	fmt.Println("c: ", c)
	fmt.Println("yC aprox.: ", yC)
	fmt.Println("error aprox.: ", errApr)

	//Newton-Raphson

	fmt.Println("Newton-Raphson method IRR")
	var fcf = []float64{-1000, 10, 10, 10, 10, 10, 10, 2000, 100}
	c, yC, errApr, k, _ := nonlineareq.NewtonRaphson(irrFunc, dxirrFunc, 0.0, 0.000001, 0.00000001, 50, fcf)
	fmt.Println("c: ", c)
	fmt.Println("yC aprox.: ", yC)
	fmt.Println("error aprox.: ", errApr)
	fmt.Println("iterations: ", k)

	fmt.Println("Steffensson method")
	c1, errSteff := nonlineareq.SteffensenAccel(irrFunc, dxirrFunc, 0, 0.000001, 0.000001, 50, fcf)
	fmt.Println("c estimated steff: ", c1)
	fmt.Println("error: ", errSteff)

	//Newton-Raphson Accelerated

	fmt.Println("Accellerated Newton-Raphson method")
	var y nonlineareq.YEqFuncx = func(x float64, params ...[]float64) float64 {
		return math.Pow(x, 3) - 3.0*x + 2
	}

	//
	var dy nonlineareq.YEqFuncx = func(x float64, params ...[]float64) float64 {
		return 3.0*math.Pow(x, 2) - 3.0
	}
	c, yC, errApr, k, errnra := nonlineareq.NewtonRaphsonAccel(y, dy, 1.2, 0.0000000000001, 0.0000000000001, 2, 50)
	fmt.Println("c: ", c)
	fmt.Println("yC aprox.: ", yC)
	fmt.Println("error aprox.: ", errApr)
	fmt.Println("iterations: ", k)
	fmt.Println("error: ", errnra)

	c, yC, errApr, k, _ = nonlineareq.NewtonRaphson(y, dy, 1.2, 0.0000000000001, 0.0000000000001, 50)
	fmt.Println("c: ", c)
	fmt.Println("yC aprox.: ", yC)
	fmt.Println("error aprox.: ", errApr)
	fmt.Println("iterations: ", k)

	fmt.Println("Steffensson method - Accel")
	c, _ = nonlineareq.SteffensenAccel(y, dy, 1.2, 0.0000000000001, 0.0000000000001, 50)
	fmt.Println("c estimated: ", c)

}

//irrFunc NPV formula for a periodic cash flow (yearly)
func irrFunc(x float64, params ...[]float64) float64 {
	lenFcf := len(params[0])
	var npv float64 = 0
	for i := 0; i < lenFcf; i++ {
		npv += params[0][i] / math.Pow(1+x, float64(i))
	}
	return npv
}

//dxirrFunc first derivative of the NPV formula for a periodic cash flow (yearly)
func dxirrFunc(x float64, params ...[]float64) float64 {
	lenFcf := len(params[0])
	var npv float64 = 0
	for i := 0; i < lenFcf; i++ {
		npv += -(float64(i) * params[0][i]) / math.Pow(1+x, float64(i)+1)
	}
	return npv
}

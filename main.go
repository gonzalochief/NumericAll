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

	//
	var irrFunc nonlineareq.YEqFuncx = func(x float64, params ...[]float64) float64 {
		lenFcf := len(params[0])
		var npv float64 = 0
		for i := 0; i < lenFcf; i++ {
			npv += params[0][i] / math.Pow(1+x, float64(i))
		}
		return npv
	}

	//
	var dxirrFunc nonlineareq.YEqFuncx = func(x float64, params ...[]float64) float64 {
		lenFcf := len(params[0])
		var npv float64 = 0
		for i := 0; i < lenFcf; i++ {
			npv += -(float64(i) * params[0][i]) / math.Pow(1+x, float64(i)+1)
		}
		return npv
	}

	fmt.Println("Newton-Raphson method IRR")
	var fcf = []float64{-1000, 10, 10, 10, 10, 10, 10, 2000, 100}
	var tv = []float64{100}
	fmt.Println("NPV: ", irrFunc(0.0, fcf, tv))
	fmt.Println("NPV: ", irrFunc(0.1177410, fcf, tv))
	c, yC, errApr, k, _ := nonlineareq.NewtonRaphson(irrFunc, dxirrFunc, 0.0, 0.000001, 0.00000001, 50, fcf)
	fmt.Println("c: ", c)
	fmt.Println("yC aprox.: ", yC)
	fmt.Println("error aprox.: ", errApr)
	fmt.Println("iterations: ", k)

}

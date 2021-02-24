package main

import (
	"fmt"
	"math"

	"github.com/gonzalochief/NumericAll/nonlineareq"
)

func main() {
	var funcToEval nonlineareq.YEqFuncx = func(x float64) float64 {
		return -4 + 4*x - (1.0/2.0)*math.Pow(x, 2)
	}
	s := funcToEval(10)
	fmt.Println("Func: ", s)
	i, pApr, errApr, relErr, pSlice, _ := nonlineareq.FixPt(funcToEval, 3.8, 10, 50)
	fmt.Println("i: ", i)
	fmt.Println("p aprox.: ", pApr)
	fmt.Println("error aprox.: ", errApr)
	fmt.Println("relative err: ", relErr)
	fmt.Println("p slice: ", pSlice)

	var funcToEval2 nonlineareq.YEqFuncx = func(x float64) float64 {
		return 0.5*x + 1.5
	}
	i, pApr, errApr, relErr, pSlice, _ = nonlineareq.FixPt(funcToEval2, 4, 10, 50)
	fmt.Println("i: ", i)
	fmt.Println("p aprox.: ", pApr)
	fmt.Println("error aprox.: ", errApr)
	fmt.Println("relative err: ", relErr)
	fmt.Println("p slice: ", pSlice)

	var funcToEval3 nonlineareq.YEqFuncx = func(x float64) float64 {
		return x*math.Sin(x) - 1
	}
	fmt.Println("evaluada: ", funcToEval3(0))
	c, yC, errApr, _ := nonlineareq.BisectBolzano(funcToEval3, 0, 2, 0.01)
	fmt.Println("c: ", c)
	fmt.Println("yC aprox.: ", yC)
	fmt.Println("error aprox.: ", errApr)

}

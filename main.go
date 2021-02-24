package main

import (
	"fmt"
	"math"
)

func main() {
	var funcToEval nonlineareq.yEqFuncx = func(x float64) float64 {
		return math.Pow(x, 5) - 3*math.Pow(x, 3) - 2*math.Pow(x, 2) + 2
	}
	s := funcToEval(1)
	fmt.Println("Func: ", s)
}

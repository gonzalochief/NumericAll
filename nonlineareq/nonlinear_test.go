package nonlineareq

import (
	"math"
	"testing"
)

type testStructureFP struct {
	TestFunctions YEqFuncx
	TestCaseName  string
	InitialEst    float64
	TestTol       int
	MaxIter       int
	ExpectedValue float64
	ExpectedIter  int
}

func TestFixedPoint(t *testing.T) {

	// Set test cases
	testCasesFP := make([]testStructureFP, 4)

	testCasesFP[0].TestFunctions = func(x float64) float64 {
		return -4 + 4*x - (1.0/2.0)*math.Pow(x, 2)
	}
	testCasesFP[0].TestCaseName = "2.1.2"
	testCasesFP[0].InitialEst = 3.8
	testCasesFP[0].TestTol = 10
	testCasesFP[0].MaxIter = 50
	testCasesFP[0].ExpectedValue = 4
	testCasesFP[0].ExpectedIter = 5

	testCasesFP[1].TestFunctions = func(x float64) float64 {
		return 1 + 2/x
	}
	testCasesFP[1].TestCaseName = "2.1.3.b"
	testCasesFP[1].InitialEst = 4
	testCasesFP[1].TestTol = 5
	testCasesFP[1].MaxIter = 50
	testCasesFP[1].ExpectedValue = 2
	testCasesFP[1].ExpectedIter = 18

	testCasesFP[2].TestFunctions = func(x float64) float64 {
		return math.Sqrt(6 + x)
	}
	testCasesFP[2].TestCaseName = "2.1.3.a"
	testCasesFP[2].InitialEst = 7
	testCasesFP[2].TestTol = 5
	testCasesFP[2].MaxIter = 50
	testCasesFP[2].ExpectedValue = 3
	testCasesFP[2].ExpectedIter = 8

	//wip
	testCasesFP[3].TestFunctions = func(x float64) float64 {
		return 0.5*x + 1.5
	}
	testCasesFP[3].TestCaseName = "Fixed Point custom 1"
	testCasesFP[3].InitialEst = 4
	testCasesFP[3].TestTol = 6
	testCasesFP[3].MaxIter = 50
	testCasesFP[3].ExpectedValue = 3.000002
	testCasesFP[3].ExpectedIter = 19

	// Test case: algorythm prformance and results
	for _, tc := range testCasesFP {
		iter, pAprox, errAprox, relErr, _, err := FixPt(tc.TestFunctions, tc.InitialEst, tc.TestTol, tc.MaxIter)
		decTol := 1.0 / math.Pow10(tc.TestTol)
		t.Logf("testing case number: %s", tc.TestCaseName)
		if err == nil {
			if math.Abs(tc.ExpectedValue-pAprox) >= decTol {
				t.Errorf("wrong values for case %s. expected aprox: %f, received: %f", tc.TestCaseName, tc.ExpectedValue, pAprox)
			}
			if iter != tc.ExpectedIter {
				t.Errorf("wrong values for case %s. expected iter: %d, received: %d", tc.TestCaseName, tc.ExpectedIter, iter)
			}
			if !((errAprox >= decTol) || (relErr >= decTol)) {
				if errAprox >= decTol {
					t.Errorf("wrong values for case %s. expected absolute error: %f, received: %f", tc.TestCaseName, decTol, errAprox)
				}
				if relErr >= decTol {
					t.Errorf("wrong values for case %s. expected relative error: %f, received: %f", tc.TestCaseName, decTol, relErr)
				}
			}
		}
		t.Logf("testing case number: %s OK", tc.TestCaseName)
		t.Logf("testing error signals case number: %s", tc.TestCaseName)
		//test case: error maximum iterations reached
		_, _, _, _, _, err = FixPt(tc.TestFunctions, tc.InitialEst, tc.TestTol, 3)
		switch err {
		case nil:
			t.Error("maximum iteration reached error not catched") // Fail to catch error. Algorithm should not converge
		}
		t.Logf("testing error signals case number: %s OK", tc.TestCaseName)
	}
}

type testStructureBisect struct {
	TestFunction    YEqFuncx
	TestCaseName    string
	TestA           float64
	TestB           float64
	TestTol         float64
	ExpectedValueC  float64
	ExpectedValueYC float64
	ExpectedErr     error
}

func TestBisect(t *testing.T) {
	// Set test cases
	testCases := make([]testStructureBisect, 4)

	testCases[0].TestFunction = func(x float64) float64 {
		return x*math.Sin(x) - 1
	}
	testCases[0].TestCaseName = "ex. 2.7"
	testCases[0].TestA = 0
	testCases[0].TestB = 2
	testCases[0].TestTol = 0.01
	testCases[0].ExpectedValueC = 1.11328125

	testCases[1].TestFunction = func(x float64) float64 {
		return 1 / (x - 2)
	}
	testCases[1].TestCaseName = "2.2.9.b"
	testCases[1].TestA = 1
	testCases[1].TestB = 7
	testCases[1].TestTol = 0.01
	testCases[1].ExpectedValueC = 1.9990234375

	testCases[2].TestFunction = func(x float64) float64 {
		return math.Tan(x)
	}

	testCases[2].TestCaseName = "2.2.10.b"
	testCases[2].TestA = 3
	testCases[2].TestB = 4
	testCases[2].TestTol = 0.01
	testCases[2].ExpectedValueC = 3.14453125

	testCases[3].TestFunction = func(x float64) float64 {
		return 1 / (x - 2)
	}
	testCases[3].TestCaseName = "2.2.9.b fail"
	testCases[3].TestA = 3
	testCases[3].TestB = 7
	testCases[3].TestTol = 0.01
	testCases[3].ExpectedValueC = 1.9990234375
	testCases[3].ExpectedErr = ErrFuncSignNotEqual

	for _, tc := range testCases {
		t.Logf("testing case number: %s", tc.TestCaseName)
		c, _, _, err := BisectBolzano(tc.TestFunction, tc.TestA, tc.TestB, tc.TestTol)
		if err == nil {
			if c != tc.ExpectedValueC {
				t.Logf("yc: %f", tc.ExpectedValueYC)
				t.Errorf("wrong estimation for case: %s. expecting: %f, receiving %f", tc.TestCaseName, tc.ExpectedValueC, c)
			}
		} else if err != tc.ExpectedErr {
			t.Error("expected error not catched") // Fail to catch error. Algorithm should not converge
		}
		t.Logf("testing case number: %s OK", tc.TestCaseName)
	}
}

type testStructRegulaFalsi struct {
	TestFunction    YEqFuncx
	TestCaseName    string
	TestA           float64
	TestB           float64
	TestTol         float64
	TestEpsilon     float64
	ExpectedValueC  float64
	ExpectedValueYC float64
	ExpectedErr     error
}

func TestRegulaFalsi(t *testing.T) {
	testCases := make([]testStructRegulaFalsi, 2)

	testCases[0].TestFunction = func(x float64) float64 {
		return x*math.Sin(x) - 1
	}
	testCases[0].TestCaseName = "ex. 2.8"
	testCases[0].TestA = 0
	testCases[0].TestB = 2
	testCases[0].TestTol = 1e-12
	testCases[0].TestEpsilon = 1e-12
	testCases[0].ExpectedValueC = 1.11415714303368
	testCases[0].ExpectedValueYC = 0

	testCases[1].TestFunction = func(x float64) float64 {
		return x*math.Sin(x) - 1
	}
	testCases[1].TestCaseName = "ex. 2.8"
	testCases[1].TestA = 0
	testCases[1].TestB = 1
	testCases[1].TestTol = 1e-12
	testCases[1].TestEpsilon = 1e-12
	testCases[1].ExpectedValueC = 1.11415714303368
	testCases[1].ExpectedValueYC = 0
	testCases[1].ExpectedErr = ErrFuncSignNotEqual

	for _, tc := range testCases {
		t.Logf("testing case number: %s", tc.TestCaseName)
		c, yC, absErr, err := RegulaFalsi(tc.TestFunction, tc.TestA, tc.TestB, tc.TestTol, tc.TestEpsilon, 50)
		if err == nil {
			if math.Abs(yC) > tc.TestEpsilon || absErr > tc.TestTol {
				t.Errorf("wrong estimation for case: %s. tolerances exceeded", tc.TestCaseName)
			}
			if tc.ExpectedValueYC != 0 {
				t.Errorf("wrong estimation for case: %s. expecting: %f, receiving %f", tc.TestCaseName, tc.ExpectedValueC, c)
			}
		} else if err != tc.ExpectedErr {
			t.Error("expected error not catched") // Fail to catch error. Algorithm should not converge
		}
		t.Logf("testing case number: %s OK", tc.TestCaseName)
		t.Logf("testing error signals case number: %s", tc.TestCaseName)
		//test case: error maximum iterations reached
		_, _, _, err = RegulaFalsi(tc.TestFunction, tc.TestA, tc.TestB, tc.TestTol, tc.TestEpsilon, 3)
		switch err {
		case nil:
			t.Error("maximum iteration reached error not catched") // Fail to catch error. Algorithm should not converge
		}
		t.Logf("testing error signals case number: %s OK", tc.TestCaseName)
	}
}

type testStructNewton struct {
	TestY            YEqFuncx
	TestDY           YEqFuncx
	TestCaseName     string
	P0               float64
	TestDelta        float64
	TestEpsilon      float64
	TestMaxIter      int
	ExpectedValueP   float64
	ExpectedValueYP  float64
	ExpectedAbsError float64
	ExpectedIter     int
}

func TestNewtonRaphson(t *testing.T) {
	testCases := make([]testStructNewton, 2)

	testCases[0].TestY = func(x float64) float64 {
		return math.Pow(x, 3) - 3*x + 2
	}
	testCases[0].TestDY = func(x float64) float64 {
		return 3*math.Pow(x, 2) - 3
	}
	testCases[0].TestCaseName = "ex. 2.14"
	testCases[0].P0 = -2.4
	testCases[0].TestDelta = 1e-12
	testCases[0].TestEpsilon = 1e-12
	testCases[0].TestMaxIter = 10
	testCases[0].ExpectedValueP = -2
	testCases[0].ExpectedValueYP = 0
	testCases[0].ExpectedAbsError = 0
	testCases[0].ExpectedIter = 4

	testCases[1].TestY = func(x float64) float64 {
		return math.Pow(x, 3) - 3*x + 2
	}
	testCases[1].TestDY = func(x float64) float64 {
		return 3*math.Pow(x, 2) - 3
	}
	testCases[1].TestCaseName = "ex. 2.15"
	testCases[1].P0 = 1.2
	testCases[1].TestDelta = 1e-12
	testCases[1].TestEpsilon = 1e-12
	testCases[1].TestMaxIter = 20
	testCases[1].ExpectedValueP = 1
	testCases[1].ExpectedValueYP = 0
	testCases[1].ExpectedAbsError = 0
	testCases[1].ExpectedIter = 18

	for _, tc := range testCases {
		t.Logf("testing case number: %s", tc.TestCaseName)
		c, yC, _, i, err := NewtonRaphson(tc.TestY, tc.TestDY, tc.P0, tc.TestDelta, tc.TestEpsilon, tc.TestMaxIter)
		if (math.Abs(tc.ExpectedValueP-c) > tc.TestDelta) && (math.Abs(tc.ExpectedValueYP) > tc.TestEpsilon) {
			if math.Abs(tc.ExpectedValueP-c) > tc.TestDelta {
				t.Errorf("wrong estimation of p for case: %s. expecting: %f, receiving %.20f", tc.TestCaseName, tc.ExpectedValueP, c)
			}
			if math.Abs(tc.ExpectedValueYP) > tc.TestEpsilon {
				t.Errorf("wrong estimation of y(p) for case: %s. expecting: %f, receiving %.20f", tc.TestCaseName, tc.ExpectedValueYP, yC)
			}
		}
		if i != tc.ExpectedIter {
			t.Errorf("wrong estimation of i for case: %s. expecting: %d, receiving %d", tc.TestCaseName, tc.ExpectedIter, i)
		}
		if err != nil {
			t.Errorf("maximum iteration reached: %d", i)
		}
		t.Logf("testing case number: %s OK", tc.TestCaseName)
		t.Logf("testing error signals case number: %s", tc.TestCaseName)
		//test case: error maximum iterations reached
		_, _, _, _, err = NewtonRaphson(tc.TestY, tc.TestDY, tc.P0, tc.TestDelta, tc.TestEpsilon, 2)
		switch err {
		case nil:
			t.Error("maximum iteration reached error not catched") // Fail to catch error. Algorithm should not converge
		}
		t.Logf("testing error signals case number: %s OK", tc.TestCaseName)
	}
}

type testStructSecant struct {
	TestY            YEqFuncx
	TestCaseName     string
	P0               float64
	P1               float64
	TestDelta        float64
	TestEpsilon      float64
	TestMaxIter      int
	ExpectedValueP   float64
	ExpectedValueYP  float64
	ExpectedAbsError float64
	ExpectedIter     int
}

func TestSecant(t *testing.T) {
	testCases := make([]testStructSecant, 2)

	testCases[0].TestY = func(x float64) float64 {
		return math.Pow(x, 3) - 3*x + 2
	}
	testCases[0].TestCaseName = "ex. 2.16"
	testCases[0].P0 = -2.6
	testCases[0].P1 = -2.4
	testCases[0].TestDelta = 1e-12
	testCases[0].TestEpsilon = 1e-12
	testCases[0].TestMaxIter = 10
	testCases[0].ExpectedValueP = -2
	testCases[0].ExpectedValueYP = 0
	testCases[0].ExpectedIter = 6

	testCases[1].TestY = func(x float64) float64 {
		return math.Pow(x, 3) - 3*x + 2
	}
	testCases[1].TestCaseName = "ex. 2.16 v2"
	testCases[1].P0 = 1.3
	testCases[1].P1 = 1.2
	testCases[1].TestDelta = 1e-12
	testCases[1].TestEpsilon = 1e-12
	testCases[1].TestMaxIter = 30
	testCases[1].ExpectedValueP = 1
	testCases[1].ExpectedValueYP = 0
	testCases[1].ExpectedIter = 26

	for _, tc := range testCases {
		t.Logf("testing case number: %s", tc.TestCaseName)
		c, yC, _, i, err := Secant(tc.TestY, tc.P0, tc.P1, tc.TestDelta, tc.TestEpsilon, tc.TestMaxIter)
		if (math.Abs(tc.ExpectedValueP-c) > tc.TestDelta) && (math.Abs(tc.ExpectedValueYP) > tc.TestEpsilon) {
			if math.Abs(tc.ExpectedValueP-c) > tc.TestDelta {
				t.Errorf("wrong estimation of p for case: %s. expecting: %f, receiving %.20f", tc.TestCaseName, tc.ExpectedValueP, c)
			}
			if math.Abs(tc.ExpectedValueYP) > tc.TestEpsilon {
				t.Errorf("wrong estimation of y(p) for case: %s. expecting: %f, receiving %.20f", tc.TestCaseName, tc.ExpectedValueYP, yC)
			}
		}
		if i != tc.ExpectedIter {
			t.Errorf("wrong estimation of i for case: %s. expecting: %d, receiving %d", tc.TestCaseName, tc.ExpectedIter, i)
		}
		if err != nil {
			t.Errorf("maximum iteration reached: %d", i)
		}

		t.Logf("testing case number: %s OK", tc.TestCaseName)
		t.Logf("testing error signals case number: %s", tc.TestCaseName)
		//test case: error maximum iterations reached
		_, _, _, _, err = Secant(tc.TestY, tc.P0, tc.P1, tc.TestDelta, tc.TestEpsilon, 2)
		switch err {
		case nil:
			t.Error("maximum iteration reached error not catched") // Fail to catch error. Algorithm should not converge
		}
		t.Logf("testing error signals case number: %s OK", tc.TestCaseName)
	}
}

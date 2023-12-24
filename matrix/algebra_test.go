type testStrMatrixSize struct {
	TestMatrixInt [][]int
	TestMatrixF64 [][]float64
	ExpectedSize  [2]int
}

func TestMatrixSize(t *testing.T) {
	testCasesInt := make([]testStrMatrixSize, 2)
	testCasesInt[0].TestMatrixInt = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	testCasesInt[0].TestMatrixF64 = [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	testCasesInt[0].ExpectedSize = [2]int{3, 3}
	testCasesInt[1].TestMatrixInt = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	testCasesInt[1].TestMatrixF64 = [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	testCasesInt[1].ExpectedSize = [2]int{2, 3}
	for _, tc := range testCasesInt {
		size := MatrixSize(tc.TestMatrixInt)
		if size != tc.ExpectedSize {
			t.Errorf("wrong matrix size estimation for int type. Expected: %v, received: %v", tc.ExpectedSize, size)
		}
		size = MatrixSize(tc.TestMatrixF64)
		if size != tc.ExpectedSize {
			t.Errorf("wrong matrix size estimation for float64 type. Expected: %v, received: %v", tc.ExpectedSize, size)
		}
	}
}

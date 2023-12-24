
// MatrixSize estimates the size of a 2D matrix
// Input:
// input is the numerical input matrix of the form [rows][column]Matrix
// Output:
// matrixSize is a vector with the size of the matrix [rows, columns]
func MatrixSize[Ord constraints.Ordered](input [][]Ord) (matrixSize [2]int) {
	matrixSize = [2]int{len(input[:][:]), len(input[0][:])}
	return
}

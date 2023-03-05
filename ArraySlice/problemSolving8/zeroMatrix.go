// Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and
// column are set to 0.

package main

import "fmt"

func zeroMatrix(matrix [][]int) [][]int {

	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))

	//traverse the matrix and record zeroes in the rows and columns slice

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	//set zeroes in the rows and columns that need to be zeroed out

	for i := 0; i < len(rows); i++ {
		if rows[i] {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < len(cols); j++ {
		if cols[j] {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}
func main() {
	givenMatrix := [][]int{
		{1, 2, 3, 4},
		{4, 0, 5, 0},
		{6, 7, 8, 9},
		{4, 6, 8, 5}}

	fmt.Println(zeroMatrix(givenMatrix))
}

// var firstRowZero bool
// 	var firstColumnZero bool

// 	for _, val := range matrix[0] {
// 		if val == 0 {
// 			firstRowZero = true
// 			break
// 		}
// 	}

// 	for i := range matrix {
// 		if matrix[i][0] == 0 {
// 			firstColumnZero = true
// 			break
// 		}
// 	}
// 	for i := 1; i < len(matrix); i++ {
// 		for j := 1; j < len(matrix[0]); j++ {
// 			if matrix[i][j] == 0 {
// 				matrix[0][j] = 0
// 				matrix[i][0] = 0
// 			}
// 		}
// 	}

// 	for i := 1; i < len(matrix); i++ {
// 		for j := 1; j < len(matrix[0]); j++ {
// 			if matrix[i][0] == 0 || matrix[0][j] == 0 {
// 				matrix[i][j] = 0
// 			}
// 		}
// 	}

// 	if firstRowZero {
// 		for j := range matrix[0] {
// 			matrix[0][j] = 0
// 		}
// 	}

// 	if firstColumnZero {
// 		for i := range matrix {
// 			matrix[i][0] = 0
// 		}
// 	}
// 	return matrix

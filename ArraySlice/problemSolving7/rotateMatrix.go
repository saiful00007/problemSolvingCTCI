//Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
//bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
//Given matrix --> Transpose Matrix --> Reverse the element of each row

package main

import "fmt"

func rotateMatrix(towDarray [][]uint8) [][]uint8 {

	length := len(towDarray)

	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			temp := towDarray[i][j]
			towDarray[i][j] = towDarray[j][i]
			towDarray[j][i] = temp
		}
	}

	for i := 0; i < length; i++ {
		start := 0
		end := length - 1

		for start < end {
			temp := towDarray[i][start]
			towDarray[i][start] = towDarray[i][end]
			towDarray[i][end] = temp
			start++
			end--
		}
	}
	return towDarray
}

func main() {
	newArray := [][]uint8{
		{1, 2, 3, 4},     //13  9   5  1
		{5, 6, 7, 8},     //14  10  6  2
		{9, 10, 11, 12},  //15  11  7  3
		{13, 14, 15, 16}} //16  12  8  4

	fmt.Println(rotateMatrix(newArray))
}

//Another way is element transfer one by one
//main logic is for row i:=0;i<matrixLength/2;i++
//and for column j:=i;j<matrixLength;j++

// package main

// import "fmt"

// func rotateMatrix(towDmatrix [][]uint8) [][]uint8 {

// 	length := len(towDmatrix)

// 	for i := 0; i < length/2; i++ {
// 		for j := i; j < length-i-1; j++ {
// 			temp := towDmatrix[i][j]
// 			towDmatrix[i][j] = towDmatrix[length-1-j][i]
// 			towDmatrix[length-1-j][i] = towDmatrix[length-i-1][length-1-j]
// 			towDmatrix[length-i-1][length-1-j] = towDmatrix[j][length-i-1]
// 			towDmatrix[j][length-i-1] = temp
// 		}
// 	}

// 	return towDmatrix
// }

// func main() {
// 	givenMatrix := [][]uint8{
// 		{1, 2, 3, 4},
// 		{5, 6, 7, 8},
// 		{9, 10, 11, 12},
// 		{13, 14, 15, 16}}

// 	fmt.Println(givenMatrix)
// 	fmt.Println(rotateMatrix(givenMatrix))
// }

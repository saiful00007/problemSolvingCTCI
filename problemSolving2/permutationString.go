//Given two strings, write a method to decide if one is a permutation of the other.
//Main logic is compare 2 string length first if the length is different return false.
//if length is same then sort both strings and return true if equal.

package main

import (
	"fmt"
	"sort"
)

// A function for converting string to rune type
func stringsToRune(str string) []rune {
	mySlice := []rune{}
	for _, char := range str {
		mySlice = append(mySlice, rune(char))
	}
	return mySlice
}

func isPermutaion(a, b string) bool {
	convertedA := stringsToRune(a)
	convertedB := stringsToRune(b)

	if len(convertedA) != len(convertedB) {
		return false
	}
	sort.SliceStable(convertedA, func(i, j int) bool {
		return convertedA[i] < convertedA[j]
	})
	sort.SliceStable(convertedB, func(i, j int) bool {
		return convertedB[i] < convertedB[j]
	})

	return string(convertedA) == string(convertedB)

}

func main() {
	fmt.Println(isPermutaion("saiful", "fulsai")) //true
	fmt.Println(isPermutaion("abcd", "abcde"))    //false
}

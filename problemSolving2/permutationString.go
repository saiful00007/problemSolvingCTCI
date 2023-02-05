package main

import (
	"fmt"
	"sort"
)

func stringsToRune(str string) []rune {
	mySlice := []rune{}
	for _, char := range str {
		mySlice = append(mySlice, rune(char))
	}
	return mySlice
}

func Permutaion(a, b string) bool {
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
	fmt.Println(Permutaion("saiful", "fulsai"))
}

// func Permutaion(a, b string) bool {
// 	// var1 := make([]string,10)
// 	// var2 := make([]string,10)

// 	var var1,var2 string
// 	if len(var1) != len(var2) {
// 		return false
// 	}
// 	sort.SliceStable(var1, func(i, j int) bool {
// 		return var1[i] < var1[j]
// 	})
// 	sort.SliceStable(var2, func(i, j int) bool {
// 		return var2[i] < var2[j]
// 	})
// 	return string(var1) == string(var2)

// }

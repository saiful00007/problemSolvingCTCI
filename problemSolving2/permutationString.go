package main

import (
	"fmt"
	"sort"
)

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

// func stringsToRunes(str string) []rune {
// 	runeArray := []rune{}
// 	for _, char := range str {
// 		runeArray = append(runeArray, rune(char))
// 	}
// 	return runeArray
// }

// func Permutation(a, b string) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	convertStrings1 := stringsToRunes(a)
// 	convertStrings2 := stringsToRunes(b)

// 	sort.SliceStable(convertStrings1, func(i, j int) bool {
// 		return convertStrings1[i] < convertStrings1[j]
// 	})
// 	sort.SliceStable(convertStrings2, func(i, j int) bool {
// 		return convertStrings2[i] < convertStrings2[j]
// 	})
// 	return string(convertStrings1) == string(convertStrings2)

// }

// func main() {
// 	fmt.Println(Permutation("abcd", "bcda"))
// }

// // func stringsToRunes(str string) []rune {
// // 	r := []rune{}
// // 	for _, v := range str {
// // 		r = append(r, rune(v))
// // 	}
// // 	return r
// // }

// // func Permutation(a, b string) bool {
// // 	if len(a) != len(b) {
// // 		return false
// // 	}
// // 	c := stringsToRunes(a)
// // 	d := stringsToRunes(b)
// // 	sort.SliceStable(c, func(i, j int) bool {
// // 		return c[i] < c[j]
// // 	})
// // 	sort.SliceStable(d, func(i, j int) bool {
// // 		return d[i] < d[j]
// // 	})
// // 	if string(c) != string(d) {
// // 		return false
// // 	}
// // 	return true
// // }

// // func main() {
// // 	fmt.Println(Permutation("abcd", "ababcd"))
// // }

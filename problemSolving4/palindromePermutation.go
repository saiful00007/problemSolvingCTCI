//here i created a function which takes a string as input and return a boolian value. ..
// First of all i create a map for holding string characters which key type is rune and value is integer...
// then i run a for loop for string input in map...
// i check if the given string have 2 odd value count if odd char more than one return false...

package main

import (
	"fmt"
)

func permutationOfPalindrome(str string) bool {
	charMap := make(map[rune]int)
	for _, char := range str {
		charMap[char]++
	}
	oddCount := 0
	for _, count := range charMap {
		if count%2 != 0 {
			oddCount++
		}
		if oddCount > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(permutationOfPalindrome("hello")) //false value
	fmt.Println(permutationOfPalindrome("radar")) //true value

}

//another function to solve this

// func PermutationOfPalindrome(str string) bool {
// 	charCount := make(map[rune]int)
// 	for _, char := range str {
// 		charCount[char]++
// 	}
// 	numOdd := 0
// 	for _, count := range charCount {
// 		if count%2 != 0 {
// 			numOdd++
// 		}
// 	}
// 	return numOdd <= 1
// }

// func main() {
// 	// testCases := []string{"racecar", "hello", "civic", "level"}
// 	// for _, testCase := range testCases {
// 	// 	fmt.Printf("%s is a permutation of a palindrome: %v\n", testCase, permutationOfPalindrome(testCase))
// 	// }
// }

//another function to solve this problem

// func checkPalindromePermutation(str string) bool {
// 	str = strings.ToLower(strings.Replace(str, " ", "", -1))
// 	oddCount := 0
// 	wordMap := make(map[string]int)
// 	for _, char := range str {
// 		if _, ok := wordMap[string(char)]; ok {
// 			wordMap[string(char)]++
// 		} else {
// 			wordMap[string(char)] = 1
// 		}
// 	}
// 	for _, char := range wordMap {
// 		if char%2 != 0 && oddCount == 1 {
// 			oddCount++
// 			if oddCount > 1 {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func main() {
// 	fmt.Println(checkPalindromePermutation("caacdef"))
// }

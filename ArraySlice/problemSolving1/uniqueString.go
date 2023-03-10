package main

import "fmt"

func main() {
	anyWord := "saiful"
	anyWord1 := "arafat mama"
	fmt.Println(isUnique(anyWord))
	fmt.Println(isUnique(anyWord1))
}

func isUnique(str string) bool {
	charMap := make(map[rune]bool)
	for _, char := range str {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}
	return true
}

// const ASCII_SIZE = 256

// func isUnique(str string) bool {
// 	if len(str) > ASCII_SIZE {
// 		return false
// 	}
// 	charArray := make([]bool, ASCII_SIZE)
// 	for _, char := range str {
// 		if charArray[char] {
// 			return false
// 		}
// 		charArray[char] = true
// 	}
// 	return true
// }

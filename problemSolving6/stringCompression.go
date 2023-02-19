// String Compression: Implement a method to perform basic string compression using the counts
// of repeated characters. For example, the string aabcccccaaa would become a2blc5a3, If the
// "compressed" string would not become smaller than the original string, your method should return
// the original string. You can assume the string has only uppercase and lowercase letters (a - z).

//main logic is we iterate the given array,if currentChar remain same currentCount=1 will be incremented
//new string will be currentChar+strconv.Itoa(currentCount)

package main

import (
	"fmt"
	"strconv"
)

func stringCompression(str string) string {
	if len(str) == 0 {
		return ""
	}
	currentChar := str[0]
	currentCount := 1
	newStr := make([]byte, 0, len(str))
	for i := 1; i < len(str); i++ {
		if str[i] == currentChar {
			currentCount++
		} else {
			newStr = append(newStr, currentChar)
			newStr = append(newStr, []byte(strconv.Itoa(currentCount))...)
			currentChar = str[i]
			currentCount = 1
		}
	}
	newStr = append(newStr, currentChar)
	newStr = append(newStr, []byte(strconv.Itoa(currentCount))...)
	if len(newStr) >= len(str) {
		return str
	}
	return string(newStr)

}

func main() {
	fmt.Println(stringCompression("aaaabbbbcd"))
	fmt.Println(stringCompression("abcd"))
	fmt.Println(stringCompression("abb"))
}

// Write a method to replace all spaces in a string with '%20'. You may assume that the string
// has sufficient space at the end to hold the additional characters, and that you are given the "true"
// length of the string. (Note: If implementing in Java, please use a character array so that you can
// perform this operation in place.)
// EXAMPLE
// Input: "Mr 3ohn Smit h 13
// Output: "Mr%203ohn%20Smith"

//declare a function which has 2 parameter,one is string another is for length and return type is string
//I have to count how many space are there in the given string
//I create a array which contain the string value as rune type,use append func for that
//main logic is newLength will be the actualLength+space*2,when we find a space we replace it by %20 for that we shoud go for a loop which is in decrease order

package main

import (
	"fmt"
)

func urlIfy(str string, actualLength int) string {
	space := 0
	newSlice := []rune{}
	for i := 0; i < actualLength; i++ {
		if string(str[i]) == " " {
			space++
		}
	}
	for _, v := range str {
		newSlice = append(newSlice, v)
	}
	newLength := space*2 + actualLength
	for i := actualLength - 1; i >= 0; i-- {
		if newSlice[i] == ' ' {
			newSlice[newLength-1] = '0'
			newSlice[newLength-2] = '2'
			newSlice[newLength-3] = '%'
			newLength = newLength - 3
		} else {
			newSlice[newLength-1] = newSlice[i]
			newLength--
		}
	}
	return string(newSlice)
}

func main() {

	fmt.Println(urlIfy("sai ful is            ", 10))
}

/*
	//using bufio new reader for full strings
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("%s\n", name)

	newName := strings.ReplaceAll(name, " ", "%20")
	fmt.Println(newName)
*/
// fmt.Println("Enter a string with space: ")
// var str string
// fmt.Scanln(&str)
// fmt.Println(str)
// newStr := strings.ReplaceAll(str, " ", "@")
// fmt.Println(newStr)

// str := "md saiful islam tuhin"
// //newStr := strings.Replace(str, " ", "%20", -1)
// fmt.Println(newStr)

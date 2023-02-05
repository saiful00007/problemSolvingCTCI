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

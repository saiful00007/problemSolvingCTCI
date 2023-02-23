// String Rotation; Assume you have a method isSubs t rin g which checks if one word is a substring
// of another. Given two strings, si and s2, write code to check if s2 is a rotation of si using only one
// call to isSubs t rin g [e.g., "wate r bottle " is a rotation oP'erbottlewat"),

// firstly we should check the length of the two given strings if not equal then return false
// if length are same then we iterate every element of the 2nd str when we found the char of first str[0] index slice the 2nd string
// if slice2+slice1==str1 then return true else return false..

package main

import (
	"fmt"
)

func stringRotation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str2); i++ {
		if str2[i] == str1[0] && str2[i:]+str2[:i] == str1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(stringRotation("waterbottle", "erbottlewat"))
}

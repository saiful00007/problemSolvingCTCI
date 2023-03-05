// There are three types of edits that can be performed on strings: insert a character,
// remove a character, or replace a character. Given two strings, write a function to check if they are
// one edit (or zero edits) away.
// EXAMPLE
// pale , ple - > true
// pales , pale - > true
// pale , bale - > true
// pale , bake - > false

//we check the length of two strings..if length is same then check how many char difference in there..
//if charDifference is more than 1 return false else return true...

package main

import "fmt"

func absulate(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func oneAwayProblem(a, b string) bool {
	if absulate(len(a)-len(b)) > 1 {
		return false
	}

	if len(a) == len(b) {
		numDifference := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				numDifference++
				if numDifference > 1 {
					return false
				}
			}
		}
		return true
	}
	if len(a) < len(b) {
		a, b = b, a
	}
	numDifference := 0
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] != b[j] {
			numDifference++
			if numDifference > 1 {
				return false
			}
			i++
		} else {
			i++
			j++
		}
	}
	return true
}

func main() {
	fmt.Println(oneAwayProblem("pale", "ple"))   //true
	fmt.Println(oneAwayProblem("pales", "pale")) //true
	fmt.Println(oneAwayProblem("pale", "bale"))  //true
	fmt.Println(oneAwayProblem("pale", "bake"))  //false
}

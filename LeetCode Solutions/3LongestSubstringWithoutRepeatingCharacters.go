/*
 * https://leetcode.com/problems/two-sum
 */
package main

import (
	"fmt"
)

var str = string("dvdf")

func main() {
	fmt.Println(bruteForce(str))
}

// Time O(n^2)
func bruteForce(s string) (result int) {
	substr := ""
	//result := 0 //for leetcode
	for i := 0; i < len(s); i++ {
		found := false
		for j := 0; j < len(substr); j++ {
			if s[j] == s[i] {
				found = true
			}
		}
		if found {
			i--
			substr = ""
		} else {
			substr = fmt.Sprint(substr + string(s[i]))
		}
		//Important this is added here, because this needs to be calculated even when no repetition occurs
		if result < len(substr) {
			result = len(substr)
		}
	}
	return result
}

// Time O(n)
/*func hashMapBased(s string) (result int) {
	seenit := make(map[byte]byte)
	substr := ""
	for i := 0; i < len(s); i++ {
		found := false
		if value, present := seenit[s[i]]; present {
			found = true
		} else {

		}
		for j := 0; j < len(substr); j++ {
			if s[j] == s[i] {
				found = true
			}
		}
		if found {
			if result < len(substr) {
				result = len(substr)
			}
			i--
			substr = ""
		} else {
			substr = fmt.Sprint(substr + string(s[i]))
		}
	}
	return result
}*/

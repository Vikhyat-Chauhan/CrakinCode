/*
 * https://leetcode.com/problems/two-sum
 */
package main

import (
	"fmt"
)

var str = string("abcabcbb")

func main() {
	fmt.Println(bruteForce(str))
}

// Time O(n^3)
func bruteForce(s string) (result int) {
	substr := ""
	for i := 0; i < len(s); i++ {
		substr = fmt.Sprint(string(s[i]))
		for j := i + 1; j < len(s); j++ {
			repeated := false
			if s[i] == s[j] {
				repeated = true
			}
			for k := 0; k < len(substr); k++ {
				if s[j] == substr[k] {
					repeated = true
				}
			}
			if !repeated {
				substr = fmt.Sprint(substr + string(s[j]))
				fmt.Println(substr)
				//if the next index is not the same
				if result < len(substr) {
					result = len(substr)
				}
			} else {
				break
			}
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

/*
 * https://leetcode.com/problems/reverse-string/
 *
 */
package main

import "fmt"

var str = []byte{'h', 'e', 'l', 'l', 'o'}

func main() {
	fmt.Println(str)
	reverseString(str)
	fmt.Println(str)
}

// Time O(n) with O(1) extra memory
func reverseString(s []byte) {
	for i := 0; i < int(len(s)/2); i++ {
		temp := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = temp
	}
}

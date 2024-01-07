/*
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/solutions/4239419/go-solution-great-explanation-and-full-description/
 */
package main

import (
	"fmt"
)

var str = string("pwwkew")

func main() {
	fmt.Println(bruteForce(str))
	fmt.Println(slidingWindowBased(str))
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
				//fmt.Println(substr)
			} else {
				break
			}
		}
		//This is outside instead of being inside because the above loop will always leave with max possible substring
		if result < len(substr) {
			result = len(substr)
		}
	}
	return result
}

// Time O(n)
func slidingWindowBased(s string) (result int) {
	//Creating a lookup map with key as char and value as last index found at
	store := make(map[uint8]int)
	var left, right int
	//Continue till we reach the last index
	for right < len(s) {
		//lets check new value in map
		r := s[right]
		//if present in map means we its repeating
		if index, present := store[r]; present {
			//move the left to the next of the repeated element using its last index in map
			if left <= index {
				left = index + 1
			}
		}
		//store the latest location of element at right
		store[r] = right
		//keep checking in every loop if len between left and right pointer has increased
		if result < (right - left + 1) {
			result = (right - left) + 1
			fmt.Println(s[left : right+1])
		}
		//increase right pointer for the next iteration
		right++
	}
	return result
}

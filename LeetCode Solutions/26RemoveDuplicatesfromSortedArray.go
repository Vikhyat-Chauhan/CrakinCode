/*
 * https://leetcode.com/problems/remove-duplicates-from-sorted-array/
 */
package main

import (
	"fmt"
)

var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

func main() {
	fmt.Println(removeDuplicates(nums))
}

// O(n) Two pointer approach
func removeDuplicates(nums []int) int {
	//First pointer this will be responsible for writing non repeated elements to the start of the array
	left := 1
	//Second pointer this will be used to visit all the elements
	for right := 1; right < len(nums); right++ {
		//Finding all locations where value changes from old index(meaning non duplicate locations)
		if nums[right] != nums[right-1] {
			//Use the left pointer to rewrite starting from the beginning
			nums[left] = nums[right]
			//increment the pointer for next iteration
			left++
		}
	}
	return left
}

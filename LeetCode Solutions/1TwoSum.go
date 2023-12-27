/*
 *
 */
package main

import "fmt"

var nums = []int{2, 7, 11, 15}
var target = 26

func main() {
	fmt.Println(bruteForce())
	fmt.Println(hashMapBased())
}

// Time O(n^2)
func bruteForce() (result []int) {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				result = append(result, i)
				result = append(result, j)
			}
		}
	}
	return result
}

// Time O(n)
func hashMapBased() (result []int) {
	seenit := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		required := target - nums[i]
		if index, present := seenit[required]; present {
			result = append(result, index)
			result = append(result, i)
			return result
		} else {
			seenit[nums[i]] = i
		}
	}
	return result
}

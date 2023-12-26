/*
 * https://www.youtube.com/watch?v=Yan0cv2cLy8&t=1s&ab_channel=NeetCode
 */
package main

import "fmt"

var nums = []int{2, 7, 11, 15}
var target = 26

func main() {
	fmt.Println(bruteForce())
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

func hashMapBased() (result []int) {
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

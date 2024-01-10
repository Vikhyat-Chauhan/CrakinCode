/*
 * https://leetcode.com/problems/two-sum
 * https://www.code-recipe.com/post/three-sum
 */
package main

import (
	"fmt"
	"slices"
	"sort"
)

var nums = []int{-1, 0, 1, 2, -1, -4}
var target = 0

func main() {
	bruteForce(nums)
	//fmt.Println(twoPointerBased(nums))
}

// Time O(n^2)
func bruteForce(nums []int) (result [][]int) {
	var tripletMap = make(map[[3]int][]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if (nums[i] + nums[j] + nums[k]) == target {
					var temp [3]int
					temp[0] = nums[i]
					temp[1] = nums[j]
					temp[2] = nums[k]
					sort.Ints(temp[:])
					tripletMap[temp] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}
	for k := range tripletMap {
		result = append(result, k[:])
		fmt.Println(k[:3])
		fmt.Println(result)
	}
	return result
}

// Time O(n)
func twoPointerBased(nums []int) (result [][]int) {
	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		if (i > 0) && (nums[i] == nums[i-1]) {
			continue
		}
		var leftI, rightI int
		leftI = i + 1
		rightI = len(nums) - 1
		for leftI < rightI {
			current := (nums[rightI] + nums[leftI] + nums[i])
			if current == target {
				result = append(result, []int{nums[i], nums[leftI], nums[rightI]})
			}
			if current < target {
				leftI++
				for nums[leftI] == nums[leftI-1] {
					leftI++
				}
			} else {
				rightI--
				for nums[rightI] == nums[rightI+1] {
					rightI--
				}
			}
		}
	}
	return result
}

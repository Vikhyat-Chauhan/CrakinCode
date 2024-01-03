/*
 * https://www.code-recipe.com/post/container-with-most-water
 */
package main

import (
	"fmt"
	"math"
)

var nums1 = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

// var nums1 = []int{1, 1}

func main() {
	fmt.Println(bruteForce(nums1))
}

// Time O(n+m)
func bruteForce(height []int) (area int) {
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			height := int(math.Min(float64(height[i]), float64(height[j])))
			breath := int(math.Abs(float64(i - j)))
			Narea := height * breath
			//fmt.Println(i, " ", j, " ", height, " ", breath)
			if area < Narea {
				area = Narea
			}
		}
	}
	return area
}

// Time O(n)
/*func hashMapBased(nums1 []int, nums2 []int) (result float64) {
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
 }*/

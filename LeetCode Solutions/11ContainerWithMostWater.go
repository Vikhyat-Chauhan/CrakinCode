/*
 * https://leetcode.com/problems/container-with-most-water/
 * https://www.code-recipe.com/post/container-with-most-water
 */
package main

import (
	"fmt"
	"math"
)

// var nums1 = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
var nums1 = []int{1, 2, 3, 4, 5, 5, 5, 7, 6}

// var nums1 = []int{1, 1}

func main() {
	fmt.Println(bruteForce(nums1))
	fmt.Println(twoPointerBased(nums1))
}

// Time O(n+m)
func bruteForce(height []int) (area int) {
	//area := 0 //for leetcode
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			heightt := int(math.Min(float64(height[i]), float64(height[j])))
			breath := int(math.Abs(float64(i - j)))
			Narea := heightt * breath
			fmt.Println(i, " ", j, " ", heightt, " ", breath, Narea)
			if area < Narea {
				area = Narea
			}
		}
	}
	return area
}

// Time O(n)
func twoPointerBased(height []int) (area int) {
	i := 0
	j := len(height) - 1
	//area := 0 //for leetcode
	for i < j {
		heightt := int(math.Min(float64(height[i]), float64(height[j])))
		breath := int(math.Abs(float64(i - j)))
		Narea := heightt * breath
		//the calculation of the area does not actually matter to the selection of pointers because it will always be chosen maximum.
		if Narea > area {
			area = Narea
		}
		//will allow us to chose all the values that will maximise the difference between pointers values
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		fmt.Println(i, " ", j, " ", heightt, " ", breath, Narea)
	}
	return area
}

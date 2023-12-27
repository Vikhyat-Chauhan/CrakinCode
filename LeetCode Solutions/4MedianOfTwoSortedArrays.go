/*
 * https://www.youtube.com/watch?v=Yan0cv2cLy8&t=1s&ab_channel=NeetCode
 */
package main

import "fmt"

var nums1 = []int{1, 2}
var nums2 = []int{3, 4}

func main() {
	fmt.Println(bruteForce(nums1, nums2))
}

// Time O(n^2)
func bruteForce(nums1 []int, nums2 []int) (result float64) {
	i := 0
	j := 0
	newarray := []int{}
	for true {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] > nums2[j] {
				newarray = append(newarray, nums1[i])
				i = i + 1
			} else {
				newarray = append(newarray, nums1[i])
				j = j + 1
			}
		} else {
			if i != len(nums1)-1 {
				newarray = append(newarray, nums1[i])
				i = i + 1
			} else {
				newarray = append(newarray, nums1[i])
				j = j + 1
			}
		}
		fmt.Println(newarray)
		return result
	}
	return result
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

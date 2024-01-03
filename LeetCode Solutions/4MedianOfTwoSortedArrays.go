/*
 * https://www.youtube.com/watch?v=Yan0cv2cLy8&t=1s&ab_channel=NeetCode
 * https://www.youtube.com/watch?v=QjrchMRAkew&t=220s&ab_channel=TechTalk
 */
package main

import "fmt"

var nums1 = []int{1, 3}
var nums2 = []int{2}

func main() {
	fmt.Println(bruteForce(nums1, nums2))
}

// Time O(n+m)
func bruteForce(nums1 []int, nums2 []int) (result float64) {
	i := 0
	j := 0
	newarray := []int{}
	//Below is for leetcode
	//var result float64
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			newarray = append(newarray, nums1[i])
			i++
		} else {
			newarray = append(newarray, nums2[j])
			j++
		}
	}
	for i < len(nums1) {
		newarray = append(newarray, nums1[i])
		i++
	}
	for j < len(nums2) {
		newarray = append(newarray, nums2[j])
		j++
	}
	fmt.Println(newarray)
	//Calculating median
	if len(newarray)%2 == 0 {
		index2 := len(newarray) / 2
		result = float64(newarray[index2-1]+newarray[index2]) / 2
	} else {
		result = float64(newarray[len(newarray)/2])
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

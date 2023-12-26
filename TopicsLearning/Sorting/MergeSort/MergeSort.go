/*
 * https://www.youtube.com/watch?v=HqPJF2L5h9U&ab_channel=AbdulBari
 * https://www.programiz.com/dsa/heap-sort
 */
package main

import "fmt"

var nums = []int{10, 20, 15, 12, 40, 25, 18}

func main() {
	fmt.Println(nums)
	fmt.Println(mergeSort(nums))
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	first := mergeSort(nums[:len(nums)/2])
	second := mergeSort(nums[len(nums)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

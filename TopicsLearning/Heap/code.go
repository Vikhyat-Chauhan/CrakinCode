/*
 * https://www.youtube.com/watch?v=HqPJF2L5h9U&ab_channel=AbdulBari
 * https://www.programiz.com/dsa/heap-sort
 */
package main

import "fmt"

var nums = []int{10, 20, 15, 12, 40, 25, 18}

func main() {
	fmt.Println(nums)
	heapSort()
	fmt.Println(nums)
}

func heapSort() {
	/*
	 * This loop takes care of getting max heap element after every heapify operation and swapping with the last element
	 * Also virtually bifercates the length of the same array and provides it to subsequent loop for limiting its size
	 */
	for j := len(nums) - 1; j >= 0; j-- {
		/*
		 * This loop performs recursive heap operation from the last length(provided by uppar loop) to the first element
		 */
		for i := j; i >= 0; i-- {
			heapify(i, j)
		}
		/*
		 * Swaps the Biggest top element from max heap we got, and replaces it with last element to maintain complete tree.
		 */
		swap := nums[j]
		nums[j] = nums[0]
		nums[0] = swap
	}
}

/*
 * Max Heapfy algorithm, this recursively starts heap operation on a complete tree starting from index provided to the max index
 */
func heapify(index int, maxindex int) {
	//Calculate children of node for comparisons later
	leftindex := (2 * index) + 1
	rightindex := (2 * index) + 2
	largest := index //lets assume patent node is largest
	// Making sure the node actually exists and then comparing subsequent values to note if leftchild is bigger than parent node
	if (leftindex <= maxindex) && (nums[leftindex] > nums[largest]) {
		largest = leftindex
	}
	// Making sure the node actually exists and then comparing subsequent values to note if rightchild is bigger than previous bigger parent or left
	if (rightindex <= maxindex) && (nums[rightindex] > nums[largest]) {
		largest = rightindex
	}else
	// Only doing swap operation once we are sure either left or right is bigger than parent node .i.e largest index is not parent
	if largest != index {
		swap := nums[largest]
		nums[largest] = nums[index]
		nums[index] = swap
		// Performing recursive operation to all the lower nodes that might be affected with the same function
		heapify(largest, maxindex)
	}
}

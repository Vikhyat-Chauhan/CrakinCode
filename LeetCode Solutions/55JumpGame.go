/*
 * https://www.youtube.com/watch?v=Yan0cv2cLy8&t=1s&ab_channel=NeetCode
 */
package main

import (
	"fmt"
	"math"
)

var nums = [5]int{2, 3, 1, 1, 4}

//var nums = [5]int{3, 2, 1, 0, 4}

/*func main() {
	success := true
	for i := 0; i < len(nums); {
		i = i + (nums[i] - 1)
		fmt.Println(i)
	}
	fmt.Println("Success : ", success)
}*/

func main() {
	jumpsLeft := int(math.Inf(-1))

	for i := 0; i < len(nums); i++ {
		if nums[i] > jumpsLeft {
			jumpsLeft = nums[i]
		}
		// have more jumps and it's not the last step
		if jumpsLeft <= 0 && i != len(nums)-1 {
			fmt.Println(false)
			break
		}

		jumpsLeft--
		fmt.Println(jumpsLeft)
	}
	fmt.Println(true)
}

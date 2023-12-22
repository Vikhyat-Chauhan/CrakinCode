// Important hint is not to be confused, with calculating the sales in the future nomatter how big the number is in the future it wont matter because the profit will remain the same even in incremental steps
// Example [1,2,3,4,5] profit = 1+1+1+1 = 4 OR profit = 5-1 = 4, meaning the algorithm does not have to worry about selling at a greater value in the future.
package main

import "fmt"

var prices = [5]int{1, 3, 5, 12, 9}

//var prices = [6]int{7, 1, 5, 3, 6, 4}
//var prices = [5]int{1, 2, 3, 4, 5}
//var prices = [5]int{7, 6, 4, 3, 1}

func main() {
	fmt.Println("Hello, World!")
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
			fmt.Println(profit)
		}
	}
	fmt.Println("Profit : ", profit)
}

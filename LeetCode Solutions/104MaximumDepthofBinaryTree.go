/*
 * https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
 * https://gobyexample.com/structs
 *
 */
package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var inputArray = []int{3, 9, 20, 0, 0, 15, 7}
var nodeHead = TreeNode{}
var target = 0

func main() {
	TreeFiller(inputArray, &nodeHead)
	//maxDepth(nums)
	fmt.Println("yellow")
}

func maxDepth(root *TreeNode) int {
	return 1
}

func TreeFiller(array []int, nodeHead *TreeNode) {
	for index := 0; true; index++ {
		leftIndex, rightIndex := (2*index)+1, (2*index)+2
	}
}

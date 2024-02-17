/*
 * https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
 * https://gobyexample.com/structs
 *
 */
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	output := make([]int, 0)

	output = append(output, left...)
	output = append(output, root.Val)
	output = append(output, right...)
	return output

}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	output := inorderTraversal(&root)
	fmt.Println(output)

}

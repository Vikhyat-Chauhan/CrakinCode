/*
 * https://leetcode.com/problems/contains-duplicate/
 *
 */
package main

// Importing fmt package for the sake of printing

type Node struct {
	value int
	next  *Node
}

type LinkedLister interface {
	new(value int)
	insert(value int)
}

func (node Node) new(value int) {
	
}

func (node Node) insert(value int) {

}

func main() {
	var linkedlist LinkedLister 
	linkedlist.new(1)
	linkedlist.insert()
}

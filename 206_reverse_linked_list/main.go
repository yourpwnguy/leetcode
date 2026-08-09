package main

import "fmt"

/*
Reverse Linked List

Approach:
We're gonna walk through the list and reverse each "Next" pointer.

Keep two pointers:
	prev -> the node we've already reversed
	curr -> the node we're currently fucking with

Before changing curr.Next, save the next node.
Otherwise, we'd overwrite the pointer and lose the rest of the list.

Then:
	next = curr.Next
	curr.Next = prev
	prev = curr
	curr = next

When curr becomes nil, prev is the new head.

Time:  O(n)
Space: O(1)
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	fmt.Println(
		reverseList(
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			}))
}

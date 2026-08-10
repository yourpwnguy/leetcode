package main

import "fmt"

/*
Linked List Cycle

Approach:
The first thought was obviously "can I just keep a map of visited nodes?"
and call it a day. But that's extra memory and we can do better.

Enter Floyd's Cycle Detection ALgorithm (Tortoise and Hare)

The idea is kinda simple:
	- slow moves one step at a time
	- fast moves two steps at a time

If a cycle exists, fast will eventually catch slow because it is moving
faster inside the loop.

If fast reaches the end of the list, there is no cycle. The linked list
just ends like a normal civilised data structure.

Time: O(n)
Space: O(1)
*/

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // cycle here

	fmt.Println(hasCycle(node1))
}

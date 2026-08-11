package main

import "fmt"

/*
Binary Search

Apporach:
The first dumb approach is scanning the whole array, but that gives O(n).
Since the array is already sorted, we can use binary search and throw away
half of the search space after every comparison.

Maintain two pointers: left and right.
Check the middle element:
 - If it matches target, return the index.
 - If middle value is smaller, target must be on the right side.
 - Otherwise, search the left side.

The array being sorted is the entire reason this works. Without ordering,
the trick is useless.

Time o(log n)
Space: O(1)
*/

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

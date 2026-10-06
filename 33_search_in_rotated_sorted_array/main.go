package main

import "fmt"

/*
Search in Rotated Sorted Array

Pattern: Rotated sorted array + O(log n) search → Modified Binary Search.

Approach:
The dumb way is checking every number one by one.
That's O(n), but the problem wants O(log n).

The array isn't completely sorted anymore because it was rotated,
but one important thing is still true:
at least ONE half of the current search range is always sorted.

So first find mid and figure out which half is sorted.

If the left half is sorted, check whether target falls between
nums[left] and nums[mid].

If it does, search the left half.
Otherwise, throw it away and search the right half.

If the left half isn't sorted, then the right half must be sorted.
Do the same check there and decide which half to keep.

Because we throw away roughly half the search space every iteration,
we still get O(log n).

The main thing to remember: when a sorted array is rotated,
think MODIFIED BINARY SEARCH — find the sorted half, then decide
whether the target can exist inside it.

Time: O(log n)
Space: O(1)
*/

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[left] < nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	result := search(nums, target)

	fmt.Println(result) // 4
}

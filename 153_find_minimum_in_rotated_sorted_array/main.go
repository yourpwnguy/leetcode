package main

import "fmt"

/*
Find Minimum in Rotated Sorted Array

Pattern: Rotated sorted array + find minimum + O(log n) → Modified Binary Search.

Approach:
The dumb way is scanning the whole array and keeping the smallest value.
That works, but it's O(n), and we need O(log n).

Since the array was originally sorted and only rotated, I can use
binary search to find where the sorted order wraps around.

Compare nums[mid] with nums[right].

If nums[mid] > nums[right], the minimum must be to the RIGHT of mid,
because mid is still inside the left sorted portion.

So:
left = mid + 1

Otherwise, nums[mid] < nums[right], which means the right side is sorted.
The minimum could be nums[mid] itself, so I need to KEEP mid in the search range.

So:
right = mid

Keep shrinking until left == right. That remaining element is the minimum.

The main thing to remember: in a rotated sorted array, compare MID with RIGHT
to decide which side contains the rotation point.

Time: O(log n)
Space: O(1)
*/

func findMin(nums []int) int {
	// Okay, I know this was originally sorted, then rotated,
	// so the smallest number is basically where that sorted order
	// got fucked up and wrapped back around to the beginning.
	//
	// I could scan the whole array and find the smallest value,
	// but that's O(n), and we need O(log n).
	//
	// So I'm gonna use binary search and keep shrinking the range
	// where the minimum can possibly be.
	left, right := 0, len(nums)-1

	// I'm gonna keep cutting the search range in half until one
	// element is left. That's our minimum.
	for left < right {
		// I'm gonna check the middle of the current search range.
		mid := left + (right-left)/2

		// Now I'll compare the middle value with the rightmost value.
		//
		// If the middle is bigger, the minimum must be to its right.
		// For example, 7 > 2 in [4, 5, 6, 7, 0, 1, 2].
		// So I'm gonna discard everything up to mid.
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			// Otherwise, the minimum is at mid or somewhere to its left.
			// I can't discard mid because it might actually be the minimum.
			// So I'll keep mid and move right to it.
			right = mid
		}
	}

	// When left == right, there's only one possible element left,
	// and that's the minimum.
	return nums[left]
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}

	result := findMin(nums)

	fmt.Println(result) // 0
}

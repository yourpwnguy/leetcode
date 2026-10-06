package main

import "fmt"

/*
Binary Search

Pattern: Sorted array + search for a value + O(log n) → Binary Search.

Approach:
The dumb way is checking every number one by one.
That's O(n), but the problem specifically wants O(log n).

Since the array is sorted, we can look at the middle element
and immediately throw away half of the search space.

If nums[mid] == target, we're done.

If nums[mid] < target, everything before mid is too small,
so move left to mid + 1.

If nums[mid] > target, everything after mid is too big,
so move right to mid - 1.

Keep doing this until we find the target or the search range becomes empty.

The main thing to remember: when the array is SORTED and we need
to search for something in O(log n), think BINARY SEARCH.

Time: O(log n)
Space: O(1)
*/

func search(nums []int, target int) int {
	// Okay, so I need to find target inside a sorted array.
	// Since the array is sorted and the problem specifically wants O(log n),
	// I can't just walk through everything one by one.
	//
	// I'm gonna use binary search, so I'll keep track of the range
	// where target could still exist.
	left, right := 0, len(nums)-1

	// Alright, while there is still some part of the array left to search,
	// I'm gonna look at the middle and decide which half I can throw away.
	for left <= right {
		// I need the middle index of my current search range.
		// Writing it this way avoids possible integer overflow
		// compared to just doing (left + right) / 2.
		mid := left + (right-left)/2

		// Okay, if the middle value is exactly what I'm looking for,
		// we're done. Just return its index.
		if nums[mid] == target {
			return mid
		}

		// If the middle value is smaller than target, then because
		// the array is sorted, everything to the left of mid is
		// also too small. So I can throw that whole half away.
		if nums[mid] < target {
			left = mid + 1
		} else {
			// Otherwise the middle value is bigger than target.
			// Since everything after mid is even bigger,
			// I don't need that half either.
			right = mid - 1
		}
	}

	// If left passed right, there is nothing left to search,
	// so target doesn't exist in the array.
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	result := search(nums, target)

	fmt.Println(result) // 4
}

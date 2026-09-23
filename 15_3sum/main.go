package main

/*
3Sum

Pattern: 3 numbers + target sum → Sort + Two Pointers.

Approach:
The dumb shit is checking every combination of 3 numbers with 3 loops.
That's O(n³), which is gonna get cooked.

Instead, sort the array first.

Then fix one number with i, and turn the rest into a 2Sum problem.
We put left after i and right at the end, then check the total sum.

If the sum is too small, move left forward because the array is sorted,
so we get a bigger number.

If the sum is too big, move right backward to get a smaller number.

If the sum is 0, we found a triplet. Add it, then move both pointers.

The annoying part is duplicates. Skip duplicate values for i, left,
and right so we don't return the same triplet multiple times.

We can also stop early if nums[i] > 0. Since the array is sorted,
everything after i is also positive, so there's no way to make 0.

The main thing to remember: when a 3Sum problem can be turned into
"fix one + find two", think SORT + TWO POINTERS.

Time: O(n²)
Space: O(1) extra space, excluding the output.
*/

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// I need somewhere to store every valid triplet I find.
	result := [][]int{}

	// I'll keep the length around because I need it multiple times.
	n := len(nums)

	// If there aren't even 3 numbers, there is no fucking way to make a triplet.
	if n < 3 {
		return result
	}

	// Sorting lets me use two pointers and also makes duplicate handling easy.
	sort.Ints(nums)

	// I'll fix one number and turn the remaining problem into a 2Sum.
	for i := range n - 2 {
		// Since the array is sorted, everything after i is also positive.
		// Three positive numbers can't possibly add up to 0, so we're fucking done.
		if nums[i] > 0 {
			return result
		}

		// If this number is the same as the previous one, I'd just find the same triplets again.
		// So skip it and let the previous iteration handle that value.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// The first number is fixed, so left starts right after it.
		// right starts at the biggest number in the array.
		left, right := i+1, n-1

		// Now I'm basically doing 2Sum between left and right.
		for left < right {
			// Add the fixed number and the two numbers we're currently checking.
			sum := nums[i] + nums[left] + nums[right]

			// If the sum is exactly 0, congratulations, we found a valid triplet.
			if sum == 0 {
				// Store the actual values, not the indexes.
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// We already used these two values, so move both pointers inward.
				left++
				right--

				// If left landed on the same value again, skip that duplicate.
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// Same shit for right. Skip repeated values so we don't return
				// the exact same triplet again.
				for left < right && nums[right] == nums[right+1] {
					right--
				}

				// The sum is too small, so I need a bigger number.
				// Since the array is sorted, moving left forward gives me that.
			} else if sum < 0 {
				left++

				// The sum is too big, so I need a smaller number.
				// Moving right backward gives me that.
			} else {
				right--
			}
		}
	}

	// We've checked every possible fixed number, so return all the triplets we found.
	return result
}

func main() {
	// Small example where we have a duplicate -1,
	// so we can actually see the duplicate handling doing its job.
	nums := []int{-1, 0, 1, 2, -1, -4}

	// Run the 3Sum solution on our input.
	result := threeSum(nums)

	// Print all valid triplets.
	fmt.Println(result)
}

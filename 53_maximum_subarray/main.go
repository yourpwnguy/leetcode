package main

import "fmt"

/*
Maximum Subarray

Pattern: Contiguous subarray + maximize/minimize sum → Kadane's Algorithm.

Approach:
The dumb shit is checking every possible subarray and calculating their sums.
That works, but we're doing way too much work for something we can solve in one pass.

Instead, we'll keep a current sum and a max sum.

For every number, we decide:
do we continue the subarray we've already built, or do we throw that shit away
and start a new subarray from the current number?

If currentSum + nums[i] is bigger, keep extending.
Otherwise, start fresh from nums[i].

Then keep track of the biggest sum we've seen so far.

Example:
For [-2, 1, -3, 4, -1, 2, 1, -5, 4], the best subarray is
[4, -1, 2, 1] with a sum of 6.

The deeper pattern:
"Should I continue what I already have, or should I start fresh here?"
→ Kadane / 1D DP.

The main thing to remember: if the running sum becomes a liability,
don't drag that shit forward. Start a new subarray.

Time: O(n)
Space: O(1)
*/

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(currentSum, maxSum)
	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

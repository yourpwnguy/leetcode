package main

import "fmt"

/*
Subarray Sum Equals K

Pattern: Subarray + exact target sum → Prefix Sum + Hash Map.

Approach:
The dumb way is checking every possible subarray and calculating its sum.
That's O(n²), and yeah, fuck doing that.

Instead, keep a running prefix sum and remember the prefix sums we've already seen.

If my current sum is 10 and k is 3, I need an old prefix sum of 7,
because 10 - 7 = 3.

So for every number, I check if I've already seen:
current sum - k

If I have, that means the numbers between that old prefix
and my current position add up to k.

The map stores:
prefix sum → how many times I've seen it

We start with prefixCount[0] = 1 because a valid subarray
can start right from index 0.

Also, we store the COUNT of each prefix sum, not just whether
it exists, because the same prefix sum can appear multiple times
and each occurrence can give us a different valid subarray.

The main thing to remember: when we need to count subarrays
with an exact sum, think PREFIX SUM + HASH MAP.

Time: O(n)
Space: O(n)
*/

func subarraySum(nums []int, k int) int {
	// Okay, so I need to count subarrays whose sum is exactly k.
	// My first thought is brute force: start at every index, keep adding numbers,
	// and whenever the sum becomes k, increase the answer.
	// But that would be O(n²), so I want to get this down to O(n).

	// The way I'm gonna do that is with prefix sums.
	// Basically, prefixSum means the sum of everything from the beginning
	// of the array up to my current position.
	//
	// So if nums is [1, 2, 3], the prefix sums are:
	// 1, 3, 6.
	//
	// I'll also keep a hashmap where the key is a prefix sum
	// and the value is how many times I've seen that prefix sum.
	prefixCount := make(map[int]int)

	// Okay, one small but important thing before I start.
	// I'm gonna say that I've already seen a prefix sum of 0 once.
	//
	// Why?
	// Because if my current prefix sum itself is k,
	// then the subarray from index 0 up to here already sums to k.
	prefixCount[0] = 1

	// This is just gonna count how many valid subarrays I've found.
	count := 0

	// And this is my running prefix sum.
	// I'll keep adding each number to it as I move through the array.
	prefixSum := 0

	// Alright, now I'm gonna walk through the array once.
	// At every position, prefixSum tells me the total sum from the beginning
	// up to this point.
	for _, num := range nums {
		// Add the current number, so now prefixSum represents
		// the sum from the start all the way up to here.
		prefixSum += num

		// Now I want to figure out whether there was some earlier prefix sum
		// that lets the part between that old position and here add up to k.
		//
		// If:
		// current prefix sum - old prefix sum = k
		//
		// then:
		// old prefix sum = current prefix sum - k
		//
		// So this is the prefix sum I need to have seen before.
		needed := prefixSum - k

		// If I've seen "needed" before, every time I saw it gives me
		// one different subarray ending at the current position
		// whose sum is exactly k.
		//
		// I'm adding the COUNT here, not just 1, because the same prefix sum
		// can appear multiple times.
		count += prefixCount[needed]

		// Okay, now that I've processed this prefix sum,
		// I need to remember it because some future subarray
		// might need this exact prefix sum.
		prefixCount[prefixSum]++
	}

	// I've now considered every possible ending position,
	// so count contains the total number of valid subarrays.
	return count
}

func main() {
	// Simple example:
	// [1, 2] adds up to 3
	// [3] also adds up to 3
	nums := []int{1, 2, 3}
	k := 3

	// Run the solution and get the number of valid subarrays.
	result := subarraySum(nums, k)

	// Should print 2.
	fmt.Println(result)
}

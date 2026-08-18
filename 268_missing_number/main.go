package main

import "fmt"

/*
Missing Number

Pattern: Range + one missing value + distinct numbers → XOR.

Approach:
The dumb approach is throwing everything into a map and checking what's missing.
Works, but we're wasting O(n) extra space for no reason.

We can use XOR because x ^ x = 0 and x ^ 0 = x.
So if we XOR all the numbers that should exist with all the numbers
actually in the array, every number that exists will cancel itself out.

The only thing left is the missing number.

We start with n because the indexes only go from 0 to n-1, but the actual
range also contains n. Then XOR each index and each value together.

The main thing to remember: when everything should appear exactly once except
one missing value, XOR can cancel out all the shit that exists in both places.

Ex:
Expected: 0 1 2 3
Actual:   3 0 1

= 0 ^ 1 ^ 2 ^ 3
^ 3 ^ 0 ^ 1

= (0 ^ 0) ^ (1 ^ 1) ^ (3 ^ 3) ^ 2
= 0 ^ 0 ^ 0 ^ 2
= 2

I'm XORing both the expected values, which I get through the index plus the
initial n, and the values actually present in the array. Since every existing
number appears twice, XOR cancels those pairs and the missing number is the
only one left.

"Why are you XORing i and num separately?"
Because i represents the numbers I'm expecting to have, while num represents
the numbers I actually have. I want both sets in the same XOR operation so
the common values cancel.

Time: O(n)
Space: O(1)
*/

func missingNumber(nums []int) int {
	n := len(nums)
	missing := n

	for i, num := range nums {
		missing ^= i
		missing ^= num
	}

	return missing
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
}

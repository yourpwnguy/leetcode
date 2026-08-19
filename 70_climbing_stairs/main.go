package main

import "fmt"

/*
Climbing Stairs

Pattern: Count ways + current depends on previous 2 → DP.

Approach:
At first this looks like some Fibonacci bullshit, but the idea is actually simple.

To reach step n, my LAST move can only be:
- 1 step → I came from n-1
- 2 steps → I came from n-2

So:
ways(n) = ways(n-1) + ways(n-2)

For example:
1 step → 1 way
2 steps → 2 ways
3 steps → 3 ways
4 steps → 5 ways

We already know the answers for 1 and 2, so start with those.
Then keep calculating the next one by adding the previous two.

No need to make a whole DP array either. We only give a fuck
about the previous 2 values, so just keep those in two variables
and slide them forward.

Time: O(n)
Space: O(1)
*/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	twoStepsBehind := 1
	oneStepBehind := 2

	for i := 3; i <= n; i++ {
		current := oneStepBehind + twoStepsBehind
		twoStepsBehind = oneStepBehind
		oneStepBehind = current
	}

	return oneStepBehind
}

func main() {
	fmt.Println(climbStairs(3))
}

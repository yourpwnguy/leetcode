package main

import "fmt"

/*
Two Sum

Approach:
Brute force would be O(n²), and we're not doing that bullshit.

Instead, keep a hash map of:
	number -> index

For every number 'v', calculate what we're missing:
	missing = target - v

Then check if we've alrady seen 'missing'.

If yes, boom. We found the pair.
If not, throw the current number into the map and keep moving.

Important:
We check the map BEFORE adding the current number, so we don't accidentally use the same element twice.

Time:  O(n)
Space: O(n)
*/

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, v := range nums {
		missing := target - v
		if j, ok := seen[missing]; ok {
			return []int{i, j}
		}
		seen[v] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

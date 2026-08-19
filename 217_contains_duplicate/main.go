package main

/*
Contains Duplicate

Pattern: Need to know if we've seen a value before → Hash Set.

Approach:
The dumb approach is comparing every number with every other number.
That works, but that's O(n²) and we're better than that shit.

Instead, keep track of the numbers we've already seen using a hash map.

For every number:
- If it's already in the map → duplicate found, return true.
- Otherwise, add it to the map and keep going.

If we reach the end without finding anything twice, return false.

Go doesn't have a built-in Set, so we use a map where the keys are the
numbers we care about.

The main thing to remember: when I need to quickly ask
"have I seen this before?", think HASH SET / HASH MAP.

Time: O(n)
Space: O(n)
*/

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

func main() {
}

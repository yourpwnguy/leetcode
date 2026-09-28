package main

import "fmt"

/*
Container With Most Water

Pattern: Two ends + maximize area → Two Pointers.

Approach:
The dumb way is checking every pair of lines and calculating
the area for each one. That's O(n²).

Instead, start with one pointer at each end of the array.
This gives us the widest possible container.

The area is:
width * height

where:
width = right - left
height = min(height[left], height[right])

The shorter line decides the height because water can't go
higher than the shorter side.

After checking the current area, move the pointer at the shorter line.
Moving the taller line doesn't help because the shorter line is
still limiting the height, while the width just gets smaller.

Keep doing this until the two pointers meet.

The main thing to remember: when two ends form an area
and we want the maximum → TWO POINTERS.

Time: O(n)
Space: O(1)
*/

func maxArea(height []int) int {
	// Okay, I need two lines that give me the biggest area.
	// So I'm gonna start with the widest possible container:
	// one pointer at the beginning and one at the end.
	left, right := 0, len(height)-1

	// I'll keep the biggest amount of water I've found so far.
	maxWater := 0

	// Alright, now I'll keep shrinking the container from both sides.
	// Every time I calculate the current area, then move one pointer inward.
	for left < right {
		// The width is just the distance between the two lines.
		width := right - left

		// The shorter line decides how tall the water can actually be.
		// Doesn't matter if the other one is fucking massive.
		waterHeight := min(height[left], height[right])

		// So area is simply width * the height we can actually use.
		currentWater := width * waterHeight

		// If this container holds more water than anything before it,
		// save it as our new answer.
		if currentWater > maxWater {
			maxWater = currentWater
		}

		// Okay, now here's the important part.
		// If the left side is shorter, moving right would only make
		// the container narrower while the short side is still there.
		// So I need to move left and hope I find a taller line.
		if height[left] < height[right] {
			left++
		} else {
			// Same idea in reverse.
			// If the right side is shorter, move right toward the middle.
			right--
		}
	}

	// We've checked all useful pairs, so return the biggest area we found.
	return maxWater
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	// Should print 49.
	result := maxArea(height)

	fmt.Println(result)
}

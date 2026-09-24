package main

import "fmt"

/*
Longest Substring Without Repeating Characters

Pattern: Longest contiguous section + maintain a condition → Sliding Window.

Approach:
The dumb approach is checking every possible substring and seeing if it has
duplicate characters. That's gonna be O(n²) or worse, and fuck that.

Instead, I'll keep a window between left and right.

Right keeps moving forward and expands the window.
If I haven't seen the current character inside my current window, we're chilling.

But if I find a duplicate, the current window is fucked.
I need to move left past the previous occurrence of that character.

The nice part is that I don't need to move left one step at a time.
I'll store the last position of every character, so when I find a duplicate
I can just jump left directly to prev + 1.

One important thing: the previous occurrence might technically exist in the map
but already be outside my current window. That's why I check prev >= left.

Every time the window is valid, calculate its length and keep the biggest one.

The main thing to remember: when the problem says
"longest/shortest substring" and there's some condition the window
needs to maintain, think SLIDING WINDOW.

Time: O(n)
Space: O(min(n, charset))
*/

func lengthOfLongestSubstring(s string) int {
	// Okay, I need to remember where I last saw each character,
	// otherwise when some character shows up again I'm gonna have no idea
	// how far I need to move left.
	lastSeen := make(map[byte]int)

	// left is basically the start of my current valid window.
	left := 0

	// This is the biggest valid window I've found so far.
	maxLen := 0

	// Right keeps moving forward and expands the window one character at a time.
	for right := range len(s) {
		// Let me see if I've already seen this character somewhere.
		if prev, ok := lastSeen[s[right]]; ok && prev >= left {
			// Ah, duplicate.
			// I can't keep both copies inside the window, so just jump left
			// past the old one instead of moving it forward one fucking step at a time.
			left = prev + 1
		}

		// This is now the most recent place where I saw this character.
		lastSeen[s[right]] = right

		// Current window is from left to right, so its length is right - left + 1.
		currentLen := right - left + 1

		// If this window is bigger than whatever bullshit I had before,
		// update the answer.
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	// We've checked the whole string, so this is the longest valid window we found.
	return maxLen
}

func main() {
	// Small example where "abc" is the longest substring without duplicates.
	s := "abcabcbb"

	// Run the sliding window solution.
	result := lengthOfLongestSubstring(s)

	// Print the maximum length we found.
	fmt.Println(result)
}

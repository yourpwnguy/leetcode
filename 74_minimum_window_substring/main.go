/*
Minimum Window Substring

Pattern: Minimum substring + satisfy character/frequency requirements → Sliding Window.

Approach:
The dumb approach is checking every possible substring and seeing if it contains
everything from t. That's gonna get expensive as fuck.

Instead, we'll use a sliding window.

First, count what characters t actually needs.
Then expand the window using right and keep track of what characters
we currently have inside it.

Once the window contains enough of every required character, the window is valid.

Now comes the important part:
because we want the MINIMUM window, we don't just keep expanding.
As soon as the window becomes valid, start moving left forward and shrink it
as much as possible while keeping it valid.

We keep the smallest valid window we've found.

The two maps are:
need   → what t requires
window → what the current window contains

And have tells us how many of the required character frequencies
are currently fully satisfied.

For example, if t = "AABC", then:
A → 2
B → 1
C → 1

Once the current window has at least those amounts, have reaches
the number of distinct requirements and the window becomes valid.

The main thing to remember:
right expands the window until it's valid,
left shrinks the valid window until it breaks.

When the problem asks for a minimum/maximum substring while maintaining
some condition, think SLIDING WINDOW.

Time: O(m + n)
Space: O(k), where k is the number of distinct characters.
*/

package main

import "fmt"

func minWindow(s string, t string) string {
	// Uhh okay, first let me handle the stupid edge case.
	// If t itself is longer than s, there's obviously no window big enough
	// to contain it, so yeah, just return an empty string and move on.
	if len(t) > len(s) {
		return ""
	}

	// Okay, so now I need to know what exactly t is asking me for.
	// Because it's not enough to just know that A exists in t.
	// If t is "AABC", I specifically need two A's, one B and one C.
	// So yeah, I'm gonna use a map and count the frequency of every character in t.
	need := make(map[byte]int)

	// So I'm just gonna walk through t once and keep adding each character
	// into the map. If I see A twice, the value for A naturally becomes 2.
	for i := range len(t) {
		need[t[i]]++
	}

	// Alright, now I know what I need.
	// But I also need to keep track of what my current sliding window actually has,
	// because at some point I need to compare "what I have" against "what I need".
	window := make(map[byte]int)

	// Now I don't really want to keep checking the entire map every fucking time
	// to figure out whether the window is valid.
	// So I'll just keep a count of how many required characters have been fully satisfied.
	have := 0

	// How many different requirements do I actually need to satisfy?
	// For "AABC", need has A, B and C, so that's 3 requirements.
	// Once have reaches 3, I know the current window has everything it needs.
	needCount := len(need)

	// Okay, now we're doing the actual sliding window.
	// I'll keep left at the beginning and let right expand the window.
	left := 0

	// I need to remember the smallest valid window I've found so far.
	// I could store the actual string every time, but that's unnecessary.
	// Just remember where it started and how long it was, and I'll slice it at the end.
	bestStart := 0
	bestLen := len(s) + 1

	// So yeah, I'm gonna write a loop over s.
	// Right is basically gonna keep walking forward and expanding my window.
	for right := range len(s) {

		// Okay, this is the character right now that I'm adding to my window.
		char := s[right]

		// So let me update the frequency of this character in the current window.
		// If I've already got two A's and I see another A, this just becomes 3.
		window[char]++

		// Now I need to ask: was this character even required by t?
		// And if it was, did I just reach the exact number of copies I needed?
		// I'm checking equality here because I only want to count this requirement once.
		if required, ok := need[char]; ok && window[char] == required {
			// Ahh okay, nice. This particular character requirement is now satisfied.
			// So I'll increase have because one more requirement is completely done.
			have++
		}

		// Okay, now if I've satisfied every requirement,
		// that means my current window is actually valid.
		// And since the problem wants the MINIMUM window,
		// this is where I stop expanding for a second and start fucking shrinking.
		for have == needCount {

			// First, before I remove anything, let me save this valid window
			// if it's smaller than the best one I've seen so far.
			currentLen := right - left + 1

			if currentLen < bestLen {
				bestLen = currentLen
				bestStart = left
			}

			// Alright, now I wanna make the window smaller from the left.
			// So whatever character is sitting at left is about to leave the window.
			leftChar := s[left]

			// Since that character is leaving, obviously I need to decrease
			// its frequency in my window.
			window[leftChar]--

			// Now here's the important part.
			// If this character was actually required by t,
			// I need to check whether removing it made me fall below the required count.
			// If I did, ahh okay, the window is no longer valid anymore.
			if required, ok := need[leftChar]; ok && window[leftChar] < required {
				// Yep, I just lost one of my required requirements.
				// So have has to go down, and that will eventually stop this shrinking loop.
				have--
			}

			// Either way, I'm done with this character at the left edge.
			// Move left forward and continue trying to make the window smaller.
			left++
		}
	}

	// Okay, we've gone through the whole string.
	// If bestLen is still this impossible value, it means we never found
	// even one valid window, so there's nothing to return.
	if bestLen == len(s)+1 {
		return ""
	}

	// Otherwise, I know exactly where my smallest window starts
	// and how long it is, so I can just slice it out and return it.
	return s[bestStart : bestStart+bestLen]
}

func main() {
	// Okay, let's use the classic example so I can actually see the shit working.
	s := "ADOBECODEBANC"
	t := "ABC"

	// Run the sliding window and grab the smallest valid substring.
	result := minWindow(s, t)

	// This should print "BANC".
	fmt.Println(result)
}

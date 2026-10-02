package main

import "fmt"

/*
Valid Anagram

Pattern: Same characters + same frequency → Frequency Count.

Approach:
The dumb way is sorting both strings and comparing them.
That works, but sorting makes it O(n log n).

Instead, count how many times each character appears.

Since the problem only has lowercase English letters, we can use
a [26]int array instead of a hashmap.

For every character in s, increase its count.
For every character in t, decrease its count.

If they're anagrams, every character appears the same number
of times in both strings, so everything cancels back to 0.

If even one count isn't 0, the strings aren't anagrams.

The main thing to remember: when two strings need to contain
the exact same characters with the exact same frequencies,
think FREQUENCY COUNT.

Time: O(n)
Space: O(1)
*/

func isAnagram(s string, t string) bool {
	// Okay, so I need to check if t is an anagram of s.
	// Basically, both strings need to have the exact same characters
	// with the exact same frequencies. The order doesn't matter.
	//
	// Like "anagram" and "nagaram" should return true because
	// they both have a:3, n:1, g:1, r:1, m:1.
	//
	// First, if their lengths are different, there's no way they
	// can have the same characters, so I can just return false.
	if len(s) != len(t) {
		return false
	}

	// Now I need some way to keep track of how many times each character
	// appears in both strings. Since the problem only has lowercase
	// English letters, I don't need a hashmap here. I can just use
	// 26 slots where index 0 is 'a', index 1 is 'b', and so on.
	var count [26]int

	// Alright, I'm gonna walk through both strings at the same time.
	// For every character in s, I'll add 1 to its count.
	// For the character at the same position in t, I'll subtract 1.
	//
	// So if both strings have the same characters with the same frequency,
	// all those +1s and -1s should cancel each other out.
	for i := range s {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	// Okay, now I just need to check whether everything actually cancelled.
	// If even one count isn't 0, that means some character appeared
	// more times in one string than the other, so they're not anagrams.
	for _, frequency := range count {
		if frequency != 0 {
			return false
		}
	}

	// Everything cancelled out, so both strings had the same characters
	// with the same frequencies. They're anagrams.
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"

	result := isAnagram(s, t)

	fmt.Println(result) // true
}

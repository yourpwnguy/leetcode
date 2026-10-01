package main

import "fmt"

/*
Group Anagrams

Pattern: String + same characters/frequencies + grouping → Hash Map.

Approach:
The dumb way is comparing every string with every other string
and checking if they're anagrams. That's gonna get ugly and slow.

Instead, I'll make a map where:
character frequency → group of anagrams

For every string, count how many times each lowercase letter appears.
Since there are only 26 lowercase letters, I can use a [26]int array.

Anagrams will always produce the exact same frequency array.

For example:
"eat" → a:1, e:1, t:1
"tea" → a:1, e:1, t:1

So both strings get the same key and end up in the same group.

I'm using the [26]int array as the map key because arrays are comparable
in Go, so they can be used as map keys.

The main thing to remember:
when strings need to be grouped based on their character counts,
think HASH MAP + FREQUENCY COUNT.

Time: O(n * k)
Space: O(n * k)

n = number of strings
k = average length of each string
*/

func groupAnagrams(strs []string) [][]string {
	// Okay, so I need to group strings that are anagrams.
	// The important thing is that the order doesn't matter.
	// "eat", "tea", and "ate" all have the same characters,
	// so they should end up in the same group.
	//
	// I'm gonna use a map to keep these groups.
	// The key will represent what characters the string has,
	// and the value will be all the strings that have that same key.
	groups := make(map[[26]int][]string)

	// Now I'm gonna loop over every string in the input slice.
	// For each string, I need to figure out what its key should be.
	for _, str := range strs {
		// I'm gonna count how many times each letter appears.
		//
		// Since the problem only has lowercase English letters,
		// I only need 26 positions:
		// 0 for 'a', 1 for 'b', ... 25 for 'z'.
		var count [26]int

		// Now I'll go through every character in this string
		// and increase the count for that character.
		for i := range str {
			// 'a' - 'a' gives 0,
			// 'b' - 'a' gives 1,
			// 'c' - 'a' gives 2, and so on.
			//
			// So this tells me exactly where this character
			// belongs in my 26-element count array.
			count[str[i]-'a']++
		}

		// Okay, now I have the character counts for this string.
		//
		// If another string is an anagram of this one,
		// it will produce the exact same count array.
		//
		// So I'll use this count as the map key and put
		// the current string into that key's group.
		groups[count] = append(groups[count], str)
	}

	// At this point the map already contains all the groups.
	// I just need to take all those groups out of the map
	// and put them into the result slice.
	result := make([][]string, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}

	// Now I have all the anagram groups.
	return result
}

func main() {
	// Classic example with three different anagram groups.
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	// Run the solution and get all the groups.
	result := groupAnagrams(strs)

	// Print the groups.
	fmt.Println(result)
}

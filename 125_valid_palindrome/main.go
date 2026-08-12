package main

import "fmt"

/*
Valid Palindrome

Pattern: Palindrome + string/array + compare both ends → Two Pointers.

Approach:
The dumb approach is cleaning the entire string, reversing it and comparing
the two strings. Works, but we're wasting memory for no reason.

Instead, we'll use two pointers:
one starts at the beginning and one starts at the end.
Move them towards each other while comparing the characters.

If a character isn't alphanumeric, skip it.
For valid characters, compare them case-insensitively.
If they don't match, the string isn't a palindrome, so return false.
If they do match, move both pointers inward and keep going.

The main thing to remember:
When we need to compare shit from both ends, think TWO POINTERS.

Time: O(n)
Space: O(1)
*/

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}

		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

package main

import "fmt"

/*
Valid Parentheses

Pattern: Nested brackets + last opened must close first → Stack.

Approach:
The first dumb thought is using a map to keep track of the brackets we've seen,
but a map doesn't give a fuck about order. We need to know what bracket was
opened most recently, so we use a stack for the LIFO part.

We'll use a map to store which closing bracket belongs to each opening bracket.
When we see an opening bracket, push it onto the stack.

When we see a closing bracket, look at the top of the stack. The map tells us
what closing bracket that opening bracket is expecting. If it doesn't match
the current character, return false. If it matches, pop it and keep going.

At the end, the stack must be empty. If it's not, some opening bracket got
left hanging around without a closing bracket.

The main thing to remember: when the last thing added needs to be handled first,
think STACK.

Time: O(n)
Space: O(n)
*/

func isValid(s string) bool {
	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', ']', '}':

			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]

			// Basically, I look at whatever bracket was opened last. The map tells me what it's supposed
			// to close with. If the current closing bracket isn't that one, the string is fucked, so I return false.
			if pairs[top] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return true
}

func main() {
	fmt.Println(isValid("([])"))
}

package main

import "fmt"

func isValidBrackets(input string) bool {
	var stack []rune

	if len(input) == 0 || len(input) >= 4096 || len(input)%2 != 0 {
		return false
	}

	bracketPairs := map[rune]rune{
		'}': '{',
		']': '[',
		'>': '<',
	}

	for _, char := range input {
		switch char {
		case '{', '[', '<':
			stack = append(stack, char)

		case '}', ']', '>':
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return false
			}

			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}

func testIsValidBracket() {
	testStrings := []string{
		"{{[<>[{{}}]]}}", // TRUE
		"[{}<>]",         // TRUE
		"]",              // FALSE
		"[>]",            // FALSE
	}

	for _, test := range testStrings {
		fmt.Printf("Input: %s, Valid: %t\n", test, isValidBrackets(test))
	}
}

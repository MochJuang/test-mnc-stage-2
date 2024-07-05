package main

import (
	"fmt"
	"strings"
)

func findMatchingStrings(n int, arrString []string) string {
	var result string
	foundMatch := false

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if strings.ToLower(arrString[i]) == strings.ToLower(arrString[j]) {
				if foundMatch {
					result += fmt.Sprintf(" %d", j+1)
				} else {
					result = fmt.Sprintf("%d %d", i+1, j+1)
					foundMatch = true
				}
			}
		}

		// if you have found string match, break the loop
		if foundMatch {
			break
		}
	}

	if !foundMatch {
		return "false"
	}
	return result
}

func testFindMatchingStrings() {
	n1 := 4
	strings1 := []string{"abcd", "acbd", "aaab", "acbd", "acbd"}
	fmt.Println(findMatchingStrings(n1, strings1)) // Output: "2 4"

	n2 := 11
	strings2 := []string{"Satu", "Sate", "Sate", "Tujuh", "Tusuk", "Tujuh", "Sate", "Bonus", "Tiga", "Puluh", "Tujuh", "Tusuk"}
	fmt.Println(findMatchingStrings(n2, strings2)) // Output: "3 5 10"

	n3 := 5
	strings3 := []string{"pisang", "goreng", "enak", "sekali", "rasanya"}
	fmt.Println(findMatchingStrings(n3, strings3)) // Output: "false"
}

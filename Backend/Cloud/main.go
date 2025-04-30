package main

import (
	"fmt"
)

func main() {

	fmt.Println(`prefix  ["flower","flow","flight"]:`, longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(`prefix  ["dog","racecar","car"]:`, longestCommonPrefix([]string{"dog", "racecar", "car"}))

}

func isValid(s string) bool {
	opposites := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	stack := make([]rune, 1, len(s))
	stack[0] = rune(s[0])

	for i := 1; i < len(s); i++ {
		if opp, ok := opposites[rune(s[i])]; ok {
			if stack[0] == opp {
				stack = stack[1:]
				continue
			}
		} else {
			stack = append(stack, rune(s[i]))
		}
	}

	return len(stack) == 0
}

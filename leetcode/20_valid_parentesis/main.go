package main

func main() {

}

func isValid(s string) bool {
	var (
		opposites = map[rune]rune{'{': '}', '(': ')', '[': ']'}
		last      int
	)

	stack := make([]rune, 1, len(s))
	stack[0] = rune(s[0])

	for i, r := range s {
		if i == 0 {
			continue
		}

		last = len(stack) - 1

		if last >= 0 && r == opposites[stack[last]] {
			stack = stack[:last]

			continue
		}

		stack = append(stack, r)
	}

	return len(stack) == 0
}

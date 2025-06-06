package main

func main() {}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	var (
		result    [][]int
		backtrack func(int, int, []int)
	)

	backtrack = func(start int, remaining int, current []int) {
		if remaining < 0 {
			return
		}

		if remaining == 0 {
			combination := make([]int, len(current))
			copy(combination, current)
			result = append(result, combination)

			return
		}

		for i := start; i < len(candidates); i++ {
			current = append(current, candidates[i])
			backtrack(i, remaining-candidates[i], current)
			current = current[:len(current)-1]
		}
	}

	backtrack(0, target, []int{})

	return result
}

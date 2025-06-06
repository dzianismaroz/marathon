package main

func main() {}

func containsDuplicate(nums []int) bool {
	var (
		duplicates = make(map[int]struct{}, len(nums))
		contains   bool
	)

	for _, n := range nums {
		if _, contains = duplicates[n]; contains {
			return true
		}

		duplicates[n] = struct{}{}
	}

	return contains
}

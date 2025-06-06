package main

func main() {}

func topKFrequent(nums []int, k int) []int {
	var (
		res    []int
		nCount = make(map[int]int, len(nums)/2)
	)

	for _, n := range nums {
		nCount[n]++
	}

	occurrences := make([][]int, len(nums)+1)

	for num, count := range nCount {
		occurrences[count] = append(occurrences[count], num)
	}

	for i := len(occurrences) - 1; i > 0; i-- {
		if len(occurrences[i]) == 0 {
			continue
		}

		for _, n := range occurrences[i] {
			if len(res) == k {
				break
			}

			res = append(res, n)
		}
	}

	return res
}

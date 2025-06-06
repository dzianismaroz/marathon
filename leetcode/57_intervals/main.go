package main

func main() {}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	start, end := newInterval[0], newInterval[1]

	var (
		merged  [][]int
		applied bool
	)

	for _, interval := range intervals {
		intervalStart, intervalEnd := interval[0], interval[1]

		if intervalEnd < start {
			merged = append(merged, interval)
			continue
		}

		if intervalStart > end {
			if !applied {
				merged = append(merged, []int{start, end})
				applied = true
			}

			merged = append(merged, interval)
			continue
		}

		start = min(start, intervalStart)
		end = max(end, intervalEnd)
	}

	if !applied {
		merged = append(merged, newInterval)
	}

	return merged
}

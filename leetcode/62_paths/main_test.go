package main

import "testing"

func TestFunc(t *testing.T) {
	tcs := []struct {
		in   [2]int
		want int
	}{
		{[2]int{3, 7}, 28},
		{[2]int{3, 2}, 3},
	}

	for tc := range tcs {
		got := uniquePaths(tcs[tc].in[0], tcs[tc].in[1])
		if got != tcs[tc].want {
			t.Errorf("got %d, want %d", got, tcs[tc].want)
		}
	}
}

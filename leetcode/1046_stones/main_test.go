package main

import "testing"

func TestStones(t *testing.T) {
	testcases := []struct {
		input []int
		want  int
	}{
		{input: []int{2, 7, 4, 1, 8, 1}, want: 1},
		{input: []int{1}, want: 1},
		{input: []int{2, 2}, want: 0},
	}

	for _, tc := range testcases {
		got := lastStoneWeight(tc.input)
		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}

package main

import "testing"

func TestLCS(t *testing.T) {
	tcs := []struct {
		in   [2]string
		want int
	}{
		{[2]string{"ABCDGH", "AEDFHR"}, 3},
		{[2]string{"AGGTAB", "GXTXAYB"}, 4},
		{[2]string{"cba", "abc"}, 1},
	}

	for tc := range tcs {
		got := longestCommonSubsequence(tcs[tc].in[0], tcs[tc].in[1])
		if got != tcs[tc].want {
			t.Errorf("got %d, want %d", got, tcs[tc].want)
		}
	}
}

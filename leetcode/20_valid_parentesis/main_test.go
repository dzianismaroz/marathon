package main

import "testing"

func TestAll(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  bool
	}{
		{
			name: "valid simple",
			in:   "()",
			out:  true,
		},

		{
			name: "other",
			in:   "()[]{}",
			out:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isValid(tc.in); got != tc.out {
				t.Errorf("failed %s got=%v want=%v", tc.name, got, tc.out)
			}
		})
	}

}

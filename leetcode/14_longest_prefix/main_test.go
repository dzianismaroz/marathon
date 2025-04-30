package main

import "testing"

func TestLtongestPrefix(t *testing.T) {
	cases := []struct {
		name string
		in   []string
		want string
	}{
		{
			name: "valid with 'fl' prefix",
			in:   []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "no matching prefix",
			in:   []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "empty",
			in:   []string{},
			want: "",
		},
		{
			name: "nil slice",
			in:   nil,
			want: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := longestCommonPrefix(tc.in); got != tc.want {
				t.Errorf("%s failed. got=%v, want=%v", tc.name, got, tc.want)
			}
		})
	}
}

func BenchmarkLongestPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = longestCommonPrefix([]string{"flower", "flow", "flight"})
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		name   string
		in     []int
		target int
		want   [][]int
	}{
		{
			name:   "empty",
			in:     []int{},
			target: 1,
			want:   [][]int{},
		},
		{
			name:   "7",
			in:     []int{2, 3, 6, 7},
			target: 7,
			want:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:   "8",
			in:     []int{2, 3, 5},
			target: 8,
			want:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(*testing.T) {
			if got := combinationSum(tc.in, tc.target); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("failed '%s' . got=%v want=%v", tc.name, got, tc.want)
			}
		})
	}
}

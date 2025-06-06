package main

import (
	"reflect"
	"testing"
)

func TestXxx(t *testing.T) {
	testcases := []struct {
		input       [][]int
		newInterval []int
		want        [][]int
	}{
		{
			input:       [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			input:       [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}

	for _, tc := range testcases {
		if got := insert(tc.input, tc.newInterval); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

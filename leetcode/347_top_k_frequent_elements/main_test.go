package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		name string
		in   []int
		k    int
		out  []int
	}{
		{
			name: "empty",
			in:   []int{},
			k:    2,
		},
		{
			name: "first",
			in:   []int{1, 1, 1, 2, 2, 3},
			k:    2,
			out:  []int{1, 2},
		},

		{
			name: "problematic",
			in:   []int{1, 2},
			k:    2,
			out:  []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := topKFrequent(tc.in, tc.k); !reflect.DeepEqual(got, tc.out) {
				t.Errorf("%s failed: got=%v, want=%v", tc.name, got, tc.out)
			}
		})
	}
}

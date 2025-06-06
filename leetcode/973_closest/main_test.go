package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	testcases := []struct {
		points [][]int
		k      int
		want   [][]int
	}{
		{
			points: [][]int{{1, 3}, {-2, 2}},
			k:      1,
			want:   [][]int{{-2, 2}},
		},
		{
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:      2,
			want:   [][]int{{3, 3}, {-2, 4}},
		},
		{
			points: [][]int{{2, 10}, {-9, -9}, {0, 8}, {-2, -2}, {8, 9}, {-10, -7}, {-5, 2}, {-4, -9}},
			k:      7,
			want:   [][]int{{-2, -2}, {-5, 2}, {0, 8}, {-4, -9}, {2, 10}, {8, 9}, {-10, -7}},
		},
	}

	for _, tc := range testcases {
		if got := kClosest(tc.points, tc.k); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

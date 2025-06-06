package main

import "testing"

func buildCyclic(t *testing.T) *ListNode {
	t.Helper()

	second := &ListNode{}
	res := &ListNode{Next: second}
	second.Next = &ListNode{Next: &ListNode{Next: second}}

	return res
}

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		name string
		in   *ListNode
		want bool
	}{
		{
			name: "empty list",
			in:   &ListNode{},
			want: false,
		},
		{
			name: "non cycle list",
			in:   &ListNode{Next: &ListNode{}},
			want: false,
		},
		{
			name: "cyclic list",
			in:   buildCyclic(t),
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := hasCycle(tc.in); got != tc.want {
				t.Errorf("failed %s, got=%v, want=%v", tc.name, got, tc.want)
			}
		})
	}
}

func BenchmarkHasCycle(b *testing.B) {
	list := buildCyclic(&testing.T{})

	b.ResetTimer()

	for b.Loop() {
		_ = hasCycle(list)
	}
}

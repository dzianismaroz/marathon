package main

import "testing"

func TestCase4(t *testing.T) {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: root}}}

	if got := detectCycle(root); got.Val != 1 {
		t.Errorf("failed cause it referess to node with value : %d ", got.Val)
	}
}

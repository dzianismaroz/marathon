package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers1(t *testing.T) {
	first := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	second := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	want := &ListNode{Val: 7, Next: &ListNode{Next: &ListNode{Val: 8}}}

	if got := addTwoNumbers(first, second); !reflect.DeepEqual(got, want) {
		t.Errorf("test case failed: want=%v; got=%v", want, got)
	}
}

func BenchmarkAddTwo(b *testing.B) {
	first := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	second := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	b.ResetTimer()

	for b.Loop() {
		_ = addTwoNumbers(first, second)
	}
}

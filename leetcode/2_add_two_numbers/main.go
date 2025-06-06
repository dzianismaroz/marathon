package main

func main() {}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		digit, acc    int
		first, second = l1, l2
		result        = &ListNode{}
		ptr           = result
	)

	for first != nil || second != nil {
		if first != nil {
			digit += first.Val
			first = first.Next
		}

		if second != nil {
			digit += second.Val
			second = second.Next
		}

		digit += acc
		acc = 0

		if digit/10 > 0 {
			digit %= 10
			acc = 1
		}

		ptr.Next = &ListNode{Val: digit}
		ptr = ptr.Next

		digit = 0
	}

	if acc != 0 {
		ptr.Next = &ListNode{Val: acc}
	}

	return result.Next
}

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head.Next, head.Next

	for slow != nil && fast != nil {
		if fast = fast.Next; fast == nil {
			return false
		}

		if slow == fast {
			return true
		}

		slow, fast = slow.Next, fast.Next
	}

	return false
}

func main() {}

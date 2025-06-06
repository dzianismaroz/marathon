package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil {
		if fast.Next == nil {
			return nil
		}

		slow, fast = slow.Next, fast.Next.Next

		if fast == slow {
			break
		}
	}

	if fast == nil {
		return nil
	}

	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

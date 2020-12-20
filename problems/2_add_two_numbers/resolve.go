package main

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	top := &ListNode{Val: 0, Next: nil}
	var cur *ListNode = top
	div := 0
	for {
		val := div
		if l1 == nil && l2 == nil && val == 0 {
			break
		}
		s1 := 0
		s2 := 0
		if l1 != nil {
			s1 = l1.Val
		}
		if l2 != nil {
			s2 = l2.Val
		}
		val += s1 + s2
		div = val / 10
		val = val % 10
		cur.Next = &ListNode{Val: val, Next: nil}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return top.Next
}

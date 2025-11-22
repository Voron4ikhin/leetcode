package internal

import "fmt"

//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
//and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.

type ListNode struct {
	Val  int
	Next *ListNode
}

func StartTwoNumbers2() {
	l111 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l11 := &ListNode{
		Val:  4,
		Next: l111,
	}
	l1 := &ListNode{
		Val:  2,
		Next: l11,
	}
	l222 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l22 := &ListNode{
		Val:  6,
		Next: l222,
	}
	l2 := &ListNode{
		Val:  5,
		Next: l22,
	}
	fmt.Println(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Ptr := l1
	l2Ptr := l2

	dummyHead := &ListNode{}
	output := dummyHead
	carry := 0
	for l1Ptr != nil || l2Ptr != nil {
		sum := carry
		if l1Ptr != nil {
			sum += l1Ptr.Val
			l1Ptr = l1Ptr.Next
		}

		if l2Ptr != nil {
			sum += l2Ptr.Val
			l2Ptr = l2Ptr.Next
		}

		carry = sum / 10
		digit := sum % 10

		output.Next = &ListNode{Val: digit}
		output = output.Next
	}

	if carry != 0 {
		output.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}

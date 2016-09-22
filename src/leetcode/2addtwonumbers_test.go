package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	for carry, cur := 0, head; l1 != nil || l2 != nil || carry != 0; cur = cur.Next {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		cur.Next = &ListNode{sum % 10, nil}
	}
	return head.Next
}

func TestAddTwoNumbers(t *testing.T) {
	assert := assert.New(t)

	assert.Equal((*ListNode)(nil), addTwoNumbers(nil, nil))

	zero := &ListNode{}
	assert.Equal(zero, addTwoNumbers(zero, zero))

	a := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	b := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	c := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}
	assert.Equal(c, addTwoNumbers(a, b))

	five := &ListNode{5, nil}
	assert.Equal(&ListNode{0, &ListNode{1, nil}}, addTwoNumbers(five, five))
}

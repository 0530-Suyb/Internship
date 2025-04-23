package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func reverse(pre *ListNode, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return reverse(head, next)
}

func reverseList2(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func removeVal(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		n1 := cur.Next
		n2 := cur.Next.Next
		cur.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		cur = n1
	}
	return dummyHead.Next
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs2(next.Next)
	next.Next = head
	return next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	fast := dummyHead
	slow := dummyHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}

func locate(head *ListNode, n int) int {
	if head.Next == nil {
		return 0
	}
	i := locate(head.Next, n) + 1
	if i == n {
		head.Next = head.Next.Next
	}
	return i
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	locate(dummyHead, n)
	return dummyHead.Next
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	l1, l2 := 0, 0
	for cur := headA; cur != nil; cur = cur.Next {
		l1++
	}
	for cur := headB; cur != nil; cur = cur.Next {
		l2++
	}
	fast, slow := headA, headB
	diff := 0
	if l1 < l2 {
		fast = headB
		slow = headA
		diff = l2 - l1
	} else {
		diff = l1 - l2
	}

	for i := 0; i < diff; i++ {
		fast = fast.Next
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

func detectCycle(head *ListNode) *ListNode {
	var ptrArr []*ListNode
	cur := head
	for cur != nil {
		for _, ptr := range ptrArr {
			if ptr == cur {
				return cur
			}
		}
		ptrArr = append(ptrArr, cur)
		cur = cur.Next
	}
	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow { // n*(y+z) = 2*(x+y) => x = (n-1)*(y+z) + z
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

func main() {
	l := Constructor()
	l.AddAtHead(1)
	fmt.Println(l.Get(0))
}

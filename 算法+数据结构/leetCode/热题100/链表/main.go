package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func isPalindrome(head *ListNode) bool {
    if head==nil || head.Next == nil {
        return true
    }
    cur := head
    arr := []int{}
    recur(cur, &arr)
    if len(arr) == 0 {
        return true
    } else {
        return false
    }
}

func recur(n *ListNode, arr *[]int) {
    if n == nil {
        return
    }
    *arr = append(*arr, n.Val)
    recur(n.Next, arr)
    if (*arr)[0] == n.Val {
        *arr = (*arr)[1:]
    }
}

func main() {
	head := &ListNode{
		Val:1, 
		Next: &ListNode{
			Val:2,
			Next: nil,
		},
	}

	fmt.Println(isPalindrome(head))
}
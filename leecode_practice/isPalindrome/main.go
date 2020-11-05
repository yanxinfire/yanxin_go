package main

import (
	"fmt"
)

func main() {
	lh := createList([]int{})
	fmt.Println(isPalindrome2(lh))
}

func createList(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode{a[len(a)-1], nil}
	for i := len(a) - 2; i >= 0; i-- {
		head = &ListNode{a[i], head}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var a []int
	for ; head != nil; head = head.Next {
		a = append(a, head.Val)
	}
	for i, v := range a[:len(a)/2] {
		if v != a[len(a)-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	fp := head
	var fcheck func(*ListNode) bool
	fcheck = func(cur *ListNode) bool {
		if cur != nil {
			if !fcheck(cur.Next) {
				return false
			}
			if cur.Val != fp.Val {
				return false
			}
			fp = fp.Next
		}
		return true
	}
	return fcheck(fp)
}

package main

import "fmt"

func main() {
	l1 := createList([]int{2, 4, 9})
	l2 := createList([]int{5, 6, 4,9})
	l3 := addTwoNumbers(l1, l2)
	for ; l3 != nil; l3 = l3.Next {
		fmt.Print(l3.Val)
	}
	fmt.Println()
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3h := &ListNode{0, nil}
	l3 := l3h
	for {
		sum:=l1.Val + l2.Val + l3.Val
		l3.Val = sum % 10
		l3.Next = &ListNode{sum/10, nil}
		switch {
		case l1.Next == nil && l2.Next == nil:
			if l3.Next.Val == 0 {
				l3.Next = nil
			}
			return l3h
		case l1.Next == nil && l2.Next != nil:
			l1 = &ListNode{0,nil}
			l2 = l2.Next
		case l1.Next != nil && l2.Next == nil:
			l2 = &ListNode{0,nil}
			l1 = l1.Next
		default:
			l1 = l1.Next
			l2 = l2.Next
		}
		l3 = l3.Next
	}
}

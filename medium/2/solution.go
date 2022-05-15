package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 ListNode
	l1.Val = 1
	var l2 ListNode
	l2.Val = 2
	l1.Next = &l2

	fmt.Println(l1, l2)
}

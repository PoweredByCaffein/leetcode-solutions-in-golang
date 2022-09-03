package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var list ListNode
	//l1.Val = 1
	//var l2 ListNode
	//l2.Val = 2
	//l1.Next = &l2

	ints := []int{1, 2, 3}

	for k, v := range ints {

		if k == 0 {
			list.Val = v
			list.Next = nil
			continue
		}

		var temp ListNode
		temp.Val = v
		temp.Next = nil

		list.Next = &temp
		
	}

	fmt.Println(list)
}

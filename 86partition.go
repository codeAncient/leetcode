package main

import (
	"fmt"
)

func main() {

	arr := []int{2, 12, 3, 11, 4, 15, 1, 7}
	list := makeLinkList(arr)
	fmt.Println(traverse(list))
	newList := partition(list, 11)
	fmt.Println(traverse(newList))
	fmt.Println("\n-----")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//86给定数分隔链表 小于x在前，大于x在后
func partition(head *ListNode, x int) *ListNode {
	//dummy1存放 <x   dummy2存放 >=x
	dummy1 := &ListNode{Val: -1}
	dummy2 := &ListNode{Val: -1}
	p := head
	p1 := dummy1
	p2 := dummy2
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		//p单立，next不再指向下一个节点， p指向下一个节点
		temp := p.Next
		p.Next = nil
		p = temp
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
func traverse(list *ListNode) string {
	num := 0
	var arr []int
	for list != nil {
		arr = append(arr, list.Val)
		num++
		list = list.Next
	}
	return fmt.Sprintf("%d nodes: %v", num, arr)
}
func makeLinkList(arr []int) *ListNode {
	dummy := &ListNode{Val: -1}
	p := dummy
	for _, val := range arr {
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	return dummy.Next
}

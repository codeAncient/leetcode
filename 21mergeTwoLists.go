package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 3, 6, 11}
	arr2 := []int{2, 6, 9, 11, 13, 15}
	list1 := makeLinkList(arr1)
	list2 := makeLinkList(arr2)
	fmt.Println(traverse(list1))
	fmt.Println(traverse(list2))
	newList := mergeTwoLists(list1, list2)
	fmt.Println(traverse(newList))
	fmt.Println("\n-----")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//21合并两个有序链表为一个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	p := dummy
	p1 := list1
	p2 := list2
	for p1 != nil && p2 != nil { //跳出循环就是其中一个越界，拷贝完毕了
		if p2.Val < p1.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}
	//谁没有越界，将剩余的第一个节点赋值给p的下一个
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	//第一个节点是虚构节点  所以返回下一个节点为头节点
	return dummy.Next
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

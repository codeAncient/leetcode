package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{38, 6, 1, 9, 34, 21, 87, 55, 97}
	list := makeLinkList(arr)

	newHead := reverseList(list)
	fmt.Println(traverse(newHead))

	arr2 := []int{38, 6, 1, 9, 34, 21, 87, 55, 97}
	list2 := makeLinkList(arr2)
	newHead2 := reverseN(list2, 3)
	fmt.Println(traverse(newHead2))

	arr3 := []int{38, 6, 1, 9, 34, 21, 87, 55, 97}
	list3 := makeLinkList(arr3)
	newHead3 := reverseBetween(list3, 2, 4)
	fmt.Println(traverse(newHead3))
}

var successor *ListNode

//反转链表的一部分，m到n
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//m=1时就是反转前n个节点
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

//反转前n个节点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 { //递归到最后一个节点，该节点的下一个节点需要记录，
		successor = head.Next
		return head
	}
	newHeader := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return newHeader
}

//反转单链表 使用递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
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

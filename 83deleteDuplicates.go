package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{38, 38, 1, 9, 9, 21, 87, 87, 87}
	list := makeLinkList(arr)
	fmt.Println(traverse(list))
	list = deleteDuplicates(list)
	fmt.Println(traverse(list))

}

//删除单链表节点中重复的元素，使用快慢两个指针，不重复元素则slow链接上新元素，slow指向当前新元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil //注意别丢了这句，后面几个重复的话要将slow的下一个节点置为nil
	return head
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

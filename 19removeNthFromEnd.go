package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//删除倒数第k个节点，p1,p2指针指向第一个节点，p1指针先往前走k步，p1,p2同时往前走n-k步，p2指向的为倒数第k个节点
func main() {
	arr := []int{3, 6, 1, 9, 34, 21, 87, 55, 97}
	list := makeLinkList(arr)
	p := findFromEnd(list, 2)
	fmt.Println(p.Val)
	list = removeNthFromEnd(list, 1)
	fmt.Println(traverse(list))

}

func findFromEnd(head *ListNode, k int) *ListNode {
	p1 := head
	p2 := head
	//p1走了k步
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	//p1 p2 同时向前走n-k步
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	//p2在n-k+1节点上 倒数第k个节点
	return p2
}

//删除倒数第n个元素，找到倒数n+1的元素  时间复杂度 O(N)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	//增加虚构头节点并不影响查找倒数第n个数是谁，加的目的在于如果共n个数，可以查倒数n+1的数就是虚拟头节点
	p := findFromEnd(dummy, n+1)
	p.Next = p.Next.Next
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

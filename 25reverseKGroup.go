package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{38, 6, 1, 9, 34, 21, 87, 55, 97}
	list := makeLinkList(arr)
	fmt.Println(traverse(list))
	newHead := reverseKGroup(list, 4)
	fmt.Println(traverse(newHead))

}

//每 k 个节点一组进行翻转
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a := head
	b := head
	for i := 0; i < k; i++ {
		if b == nil { //不够k个元素 直接返回头结点
			fmt.Printf("head:%d", head.Val)
			return head
		}
		b = b.Next
	}
	if b == nil {
		fmt.Println("cur-b:nil")
	} else {
		fmt.Printf("cur-b:%d\n", b.Val)
	}

	//循环后b节点为第k+1个元素
	newHead := reverseAB(a, b)
	a.Next = reverseKGroup(b, k)
	fmt.Printf("newHead:%d\n", newHead.Val)
	return newHead
}

//反转单链表 使用循环 每一个指向前一个
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode
	cur := head

	for cur != nil {
		next = cur.Next
		cur.Next = pre //指向前一个
		pre = cur      //pre指向当前节点
		cur = next     //当前节点向下移动
	}
	return pre
}

//反转单链表[a,b)区间
func reverseAB(a *ListNode, b *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode
	cur := a

	for cur != b {
		next = cur.Next
		cur.Next = pre //指向前一个
		pre = cur      //pre指向当前节点
		cur = next     //当前节点向下移动
	}
	return pre
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

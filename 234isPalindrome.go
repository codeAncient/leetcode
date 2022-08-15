package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{1, 2, 3, 4}
	list := makeLinkList(arr)
	fmt.Println(traverse(list))
	fmt.Println(isPalindrome(list))
	fmt.Println(traverse(list))

}

//判断是否回文，快慢指针，得到中间点，中间位置之后的点进行链表反转，这样左右比对判断
func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	slowPre := head
	for fast != nil && fast.Next != nil { //中心奇数fast不为nil
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
		slowPre = slowPre.Next
	}
	//fmt.Printf("fast:%v slowPre:%v\n  slow:%v\n", fast, slowPre, slow)
	left := head
	//fmt.Printf("noreverse:%v\n", traverse(slow))
	right := reverseList(slow)
	//fmt.Printf("reverse:%v\n", traverse(right))
	newHead := right
	for right != nil {
		if right.Val != left.Val {
			slowPre.Next = reverseList(newHead)
			return false
		}
		right = right.Next
		left = left.Next
	}
	slowPre.Next = reverseList(newHead)
	return true
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

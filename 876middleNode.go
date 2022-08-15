package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{38, 6, 1, 9, 34, 21, 87, 55, 97}
	list := makeLinkList(arr)
	//求中间节点后的链表
	p := middleNodeBefore(list)
	fmt.Println(traverse(p))
	//是否有环
	head, tail := makeLinkListHasTail(arr)
	fmt.Println(hasCycle(head))
	//连成环
	tail.Next = head
	fmt.Println(hasCycle(head))
	arr2 := []int{37, 16, 12, 19}
	head2, tail2 := makeLinkListHasTail(arr2)
	//第二个链表的尾部连上上面环的head
	tail2.Next = head
	//测试环起点，应该为head的val值
	cycleBeg := detectCycle(head2)
	fmt.Println(cycleBeg.Val)
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil { //能跳出循环肯定是无环
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { //指针相同，速度不同相遇，快的比慢的多跑1圈
			return true
		}
	}
	return false
}

//链表中包含环，求环的起始位置
//如果有环，速度不同总会相遇。相遇时slow走了k步，fast走了2k，假设相遇距离环起点m，fast再走k步到达相遇点，走k-m则到环起点
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

//876查找中间节点 偶数返回后面那个，奇数则中间1个
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil { //next为nil代表来到了边界 本身为nil是偶数个数时
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//查找中间节点 偶数返回前面那个，奇数则中间1个
func middleNodeBefore(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil { //next为nil代表来到了边界 本身为nil是偶数个数时

		fast = fast.Next.Next
		if fast != nil {
			slow = slow.Next
		}
	}
	return slow
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

func makeLinkListHasTail(arr []int) (*ListNode, *ListNode) {
	dummy := &ListNode{Val: -1}
	p := dummy
	for _, val := range arr {
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	return dummy.Next, p
}

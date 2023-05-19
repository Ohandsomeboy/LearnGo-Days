package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}

// 在这段代码中，dummyHead 是一个指向 ListNode 结构体的指针。
// 通过 &ListNode{Next: head} 创建了一个新的 ListNode 对象，
// 并使用 Next: head 的方式将新节点的 Next 字段设置为原链表的头节点 head。
func deleteMiddle(head *ListNode) *ListNode {
	// 创建一个虚拟头节点，指向链表的头部
	dummyHead := &ListNode{Next: head}
	// 使用快慢指针来定位中间节点的前一个节点
	slow, fast := dummyHead, dummyHead

	// 使用快慢指针遍历链表，每次快指针移动两步，慢指针移动一步
	// 当快指针到达链表末尾时，慢指针指向链表的中间节点的前一个节点
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 将中间节点的前一个节点的Next指针指向中间节点的后一个节点
	slow.Next = slow.Next.Next

	// 返回修改后的链表头节点
	return dummyHead.Next
}

func main() {
	var head = new(ListNode)
	head.Val = 1
	var Node1 = new(ListNode)
	head.Next = Node1
	Node1.Val = 2
	var Node2 = new(ListNode)
	Node1.Next = Node2
	Node2.Val = 3
	var Node3 = new(ListNode)
	Node2.Next = Node3
	Node3.Val = 4
	var Node4 = new(ListNode)
	Node3.Next = Node4
	Node4.Val = 5

	printList(deleteMiddle(head))
}

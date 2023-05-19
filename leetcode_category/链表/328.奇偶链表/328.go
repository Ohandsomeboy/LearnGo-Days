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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil { // 如果链表为空，则直接返回
		return head
	}
	evenHead := head.Next                 // 保存偶数节点链表的头结点
	odd := head                           // 初始化奇数节点指针为头结点
	even := evenHead                      // 初始化偶数节点指针为偶数链表的头结点
	for even != nil && even.Next != nil { // 循环条件：偶数节点和其后的节点都存在
		odd.Next = even.Next // 将 奇数节点的下一个节点 指向 偶数节点的下一个节点
		odd = odd.Next       // 更新奇数节点指针
		even.Next = odd.Next // 将 偶数节点的下一个节点 指向 奇数节点的下一个节点
		even = even.Next     // 更新偶数节点指针
	}
	odd.Next = evenHead // 将奇数节点的尾部连接到偶数链表的头部
	return head         // 返回原始链表的头结点
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

	printList(oddEvenList(head))
}

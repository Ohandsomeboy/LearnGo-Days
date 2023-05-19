package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList_head(head *ListNode) *ListNode {
	var newHead *ListNode // 创建一个变量来存储反转后链表的新头结点
	for head != nil {     // 遍历原链表直到到达尾部
		next := head.Next   // 将下一个节点存储在临时变量next中
		head.Next = newHead // 将当前节点的Next指针指向新头结点newHead
		newHead = head      // 更新新头结点为当前节点
		head = next         // 移动到原链表的下一个节点
	}
	return newHead // 返回反转后链表的新头结点
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}

func reverseList_2index(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 递归反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
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

	//var head2 = new(ListNode)
	//head2.Val = 4
	//var Node3 = new(ListNode)
	//head2.Next = Node3
	//Node3.Val = 5
	//var Node4 = new(ListNode)
	//Node3.Next = Node4
	//Node4.Val = 6

	printList(reverseList_head(head))
}

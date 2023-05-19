package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 至于本题并不难，只需要判断两个链表头部元素的大小，
// 把小的那个链表节点放到 结果链表 的结尾即可。
func reverseList_head(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
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

	printList(reverseList(head))
}

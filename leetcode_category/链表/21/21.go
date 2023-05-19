package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 分享我做链表题目的一些套路。
//
// 哑节点
// 创建 哑节点 作为 结果链表 的开头，返回结果是这个节点的下一个位置。目的是：
// 在未遍历之前，我们不知道构建的结果中，开头元素到底是 l1 还是 l2, 为了让代码整齐，创建哑节点。
//
// 使用 move 游标
// 哑节点标记了 结果链表 的开头，因此是不能移动的。为了把两个链表 merge 的结果放到结果链表的最后，
// 就需要使用一个 move 游标指向 结果链表 的最后一个元素。初始时，move 指向 哑节点，
// 之后随着结果链表的增加而不停地向后移动，始终保持其指向 结果链表 的最后一个元素。
//
// while 遍历两个元素
// 涉及到两个元素的遍历题，使用 while l1 and/or l2 的方式。
// 即两个元素都没有遍历完或者至少有一个没遍历完，具体使用 and 还是 or 要根据场景进行选择。
// 这类常见的题目有：
//
// 合并两个链表
// 两数相加/两个链表表示的数相加
// 没用完的元素仍需拼接
// 当 while 循环结束之后，l1 和 l2 至少遍历完了一个，另一个链表可能没有用完，因此需要拼接到 结果链表 的结尾。
// 合并链表 或者 两数相加 都要记得这个问题。

// 至于本题并不难，只需要判断两个链表头部元素的大小，
// 把小的那个链表节点放到 结果链表 的结尾即可。
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// dummyHead是一个哑结点，用于方便操作链表的头节点
	dummyHead := &ListNode{}
	// 定义指针p指向dummyHead，用于遍历并操作新链表
	p := dummyHead
	// 循环遍历两个链表，将它们中较小的元素依次添加到新链表的尾部
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			p.Next = list2
			list2 = list2.Next
		} else {
			p.Next = list1
			list1 = list1.Next
		}
		p = p.Next
	}
	// 如果list1还有剩余元素，则直接将剩余部分连接到新链表的尾部
	if list1 != nil {
		p.Next = list1
	}
	// 如果list2还有剩余元素，则直接将剩余部分连接到新链表的尾部
	if list2 != nil {
		p.Next = list2
	}
	// 返回新链表的头节点
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

	var head2 = new(ListNode)
	head2.Val = 4
	var Node3 = new(ListNode)
	head2.Next = Node3
	Node3.Val = 5
	var Node4 = new(ListNode)
	Node3.Next = Node4
	Node4.Val = 6

	fmt.Println(mergeTwoLists(head, head2))
}

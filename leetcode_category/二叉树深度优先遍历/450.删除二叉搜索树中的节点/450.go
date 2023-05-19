package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// deleteNode 从二叉搜索树中删除一个节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	// 如果根节点为空，返回 nil
	case root == nil:
		return nil
	// 如果要删除的值比根节点小，递归地在左子树中删除
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	// 如果要删除的值比根节点大，递归地在右子树中删除
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	// 如果要删除的节点只有一个子节点或者没有子节点，直接返回该子节点或者 nil
	case root.Left == nil || root.Right == nil:
		if root.Left != nil {
			return root.Left
		}
		return root.Right
	// 如果要删除的节点有两个子节点，找到右子树中最小的节点作为后继，将其值赋给当前节点，然后递归地在右子树中删除后继。
	default:
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		successor.Right = deleteNode(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
	return root
}

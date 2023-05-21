package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	//postorder 是一个函数变量，它的类型是 func(node *TreeNode)。
	//这意味着 postorder 可以存储一个函数，
	//该函数接收一个 *TreeNode 类型的参数，没有返回值。
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return res
}

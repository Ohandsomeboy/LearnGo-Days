package main

import "reflect"

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	//r1 := []int{}
	//r2 := []int{}
	return reflect.DeepEqual(dfs(root1), dfs(root2))
}

func dfs(root *TreeNode) []int {
	res := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if node.Left == nil && node.Right == nil {
			res = append(res, node.val)
		}
		inorder(node.Right)
	}
	inorder(root)
	return res
}

func createTree() *TreeNode {
	root := &TreeNode{val: 1}
	root.Left = &TreeNode{val: 2}
	root.Right = &TreeNode{val: 3}
	root.Left.Left = &TreeNode{val: 4}
	root.Left.Right = &TreeNode{val: 5}
	root.Right.Left = &TreeNode{val: 6}
	root.Right.Right = &TreeNode{val: 7}
	return root
}

func createTree2() *TreeNode {
	root := &TreeNode{val: 1}
	root.Left = &TreeNode{val: 2}
	root.Right = &TreeNode{val: 3}
	root.Left.Left = &TreeNode{val: 4}
	root.Left.Right = &TreeNode{val: 5}
	root.Right.Left = &TreeNode{val: 6}
	root.Right.Right = &TreeNode{val: 7}
	return root
}

func main() {
	println(leafSimilar(createTree(), createTree2()))
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(arr [][]int, row, col int) *TreeNode {
	if row == len(arr) {
		return nil
	}

	val := arr[row][col]

	left := buildTree(arr, row+1, col)
	right := buildTree(arr, row+1, col+1)

	return &TreeNode{Val: val, Left: left, Right: right}
}

func maxSumTree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftSum := maxSumTree(node.Left)
	rightSum := maxSumTree(node.Right)

	maxChildSum := leftSum
	if rightSum > leftSum {
		maxChildSum = rightSum
	}

	maxSum := node.Val
	if node.Val+maxChildSum > maxSum {
		maxSum = node.Val + maxChildSum
	}

	return maxSum
}

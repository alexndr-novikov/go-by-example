package main

import (
	"testing"
)

// run as  go test -bench=. -benchtime 5s -benchmem
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return recursiveValidate(root, nil, nil)
}

func recursiveValidate(node, min, max *TreeNode) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Val <= min.Val {
		return false
	}

	if max != nil && node.Val >= max.Val {
		return false
	}

	return recursiveValidate(node.Left, min, node) && recursiveValidate(node.Right, node, max)
}

func isValidBSTInOptimal(root *TreeNode) bool {
	res := extractInorder(root)
	minV := res[0]
	for i := 1; i < len(res); i++ {
		if res[i] <= minV {
			return false
		}
		minV = res[i]
	}
	return true
}

func extractInorder(node *TreeNode) []int {
	var res []int
	if node == nil {
		return res
	}
	res = append(extractInorder(node.Left), res...)
	res = append(res, node.Val)
	res = append(res, extractInorder(node.Right)...)
	return res
}

func BenchmarkOptimal(b *testing.B) {
	tree1 := &TreeNode{
		Val: 120,
		Left: &TreeNode{
			Val: 70,
			Left: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val:   20,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   55,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val:   75,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   110,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 140,
			Left: &TreeNode{
				Val: 130,
				Left: &TreeNode{
					Val:   119,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   135,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val: 160,
				Left: &TreeNode{
					Val:   150,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   200,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	for i := 0; i < b.N; i++ {
		isValidBST(tree1)
		isValidBST(tree2)
	}
}

func BenchmarkInOptimal(b *testing.B) {
	tree1 := &TreeNode{
		Val: 120,
		Left: &TreeNode{
			Val: 70,
			Left: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val:   20,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   55,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val: 100,
				Left: &TreeNode{
					Val:   75,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   110,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 140,
			Left: &TreeNode{
				Val: 130,
				Left: &TreeNode{
					Val:   119,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   135,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val: 160,
				Left: &TreeNode{
					Val:   150,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   200,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	for i := 0; i < b.N; i++ {
		isValidBSTInOptimal(tree1)
		isValidBSTInOptimal(tree2)
	}
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}
func recur(A *TreeNode, B *TreeNode) bool {
	// B 为空 表示已经遍历完成 return true
	if B == nil {
		return true
	}
	// A 为空 表示已经比对完成 return false
	if A == nil {
		return false
	}
	// A B 不相等表示 不匹配 return false
	if A.Val != B.Val {
		return false
	}
	// 其他接着执行
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}

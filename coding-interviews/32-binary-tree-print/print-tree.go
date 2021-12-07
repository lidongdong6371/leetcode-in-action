/**
## 题目
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]


提示：

节点总数 <= 1000


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ans := []int{}
	for len(queue) != 0 {
		cRoot := queue[0]
		ans = append(ans, cRoot.Val)
		if cRoot.Left != nil {
			queue = append(queue, cRoot.Left)
		}

		if cRoot.Right != nil {
			queue = append(queue, cRoot.Right)
		}

		queue = queue[1:]
	}

	return ans
}
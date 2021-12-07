/*
	## 题目 反转链表

	定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。



示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL


限制：

0 <= 节点个数 <= 5000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof
*/

/*
	## 题解&思路

	使用双指针，cur, pre
	cur 初始为当前节点， pre为cur 前一个节点，初始为 null
	cur 和 pre 通过向下迭代，变更原有的指向，知道 cur 为 null
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

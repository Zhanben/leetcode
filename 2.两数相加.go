//package main

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start

//Definition for singly-linked list.
// //

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	flag := 0
	for l1 != nil || l2 != nil || flag != 0 {
		if l1 != nil {
			flag += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			flag += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{
			Val:  flag % 10,
			Next: nil,
		}
		cur = cur.Next
		flag = flag/10

	}
	return dummy.Next

}

// @lc code=end

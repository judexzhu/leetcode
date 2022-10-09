/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 *
 * https://leetcode.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (49.99%)
 * Likes:    7335
 * Dislikes: 255
 * Total Accepted:    578.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3,4]'
 *
 * You are given the head of a singly linked-list. The list can be represented
 * as:
 * 
 * 
 * L0 → L1 → … → Ln - 1 → Ln
 * 
 * 
 * Reorder the list to be on the following form:
 * 
 * 
 * L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 * 
 * 
 * You may not modify the values in the list's nodes. Only nodes themselves may
 * be changed.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: head = [1,2,3,4]
 * Output: [1,4,2,3]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: head = [1,2,3,4,5]
 * Output: [1,5,2,4,3]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in the list is in the range [1, 5 * 10^4].
 * 1 <= Node.val <= 1000
 * 
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reorderList(head *ListNode)  {
	// create an empty array/slice  
	nodes := []*ListNode{}

	// add all nodes to the array
	for node := head; node != nil; node = node.Next{
		nodes = append(nodes, node)
	}

	//use two pointer
	//move two point from the every left and right to middle
    for i,j := 0,len(nodes)-1;i<j;i,j=i+1,j-1 {
        next := nodes[i].Next
		nodes[i].Next = nodes[j]
        nodes[j].Next = next
	}
	// Dont forget to set the mid, very end of the new listNode point to nil
	nodes[len(nodes)/2].Next = nil
}
// @lc code=end


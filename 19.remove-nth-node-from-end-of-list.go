/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (39.05%)
 * Likes:    13337
 * Dislikes: 549
 * Total Accepted:    1.7M
 * Total Submissions: 4.3M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given the head of a linked list, remove the n^th node from the end of the
 * list and return its head.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: head = [1,2,3,4,5], n = 2
 * Output: [1,2,3,5]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: head = [1], n = 1
 * Output: []
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: head = [1,2], n = 1
 * Output: [1]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in the list is sz.
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
 * 
 * 
 * 
 * Follow up: Could you do this in one pass?
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	//Use two pointers
	// set two pointers to the head
	slow := head
	fast := head 

	// move fast to the n position to the slow pointer
	for i:=0; i<n; i++ {
		fast = fast.Next
	}
	// now if fast is nil (out of the listNode)
	// that means which need to remove is the head
	if fast == nil {
		return head.Next
	}
	// keep moveing the slowpoint and fast point
	// when fast pointer is to the very end of the listNodes
	// slow.Next is the node need to be removed
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next
	}
   // what we need to now is let slow skip one next
   // return the head 
	slow.Next = slow.Next.Next

	return head
    
}
// @lc code=end


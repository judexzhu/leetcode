/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (61.35%)
 * Likes:    14409
 * Dislikes: 1284
 * Total Accepted:    2.5M
 * Total Submissions: 4.1M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * You are given the heads of two sorted linked lists list1 and list2.
 * 
 * Merge the two lists in a one sorted list. The list should be made by
 * splicing together the nodes of the first two lists.
 * 
 * Return the head of the merged linked list.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: list1 = [1,2,4], list2 = [1,3,4]
 * Output: [1,1,2,3,4,4]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: list1 = [], list2 = []
 * Output: []
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: list1 = [], list2 = [0]
 * Output: [0]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in both lists is in the range [0, 50].
 * -100 <= Node.val <= 100
 * Both list1 and list2 are sorted in non-decreasing order.
 * 
 * 
 */
package leetcode
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    // create a emtpy dummpy node 
    dummy := &ListNode{}
    tail := dummy
    // if list1 and list2 are not emtpty(nil)
    for list1 != nil && list2 != nil{
        // if current list1 node value is smaller
        if list1.Val < list2.Val{
            // link cuurent list1 node to the dummy node next
            tail.Next = list1
            // move current list1 node to the list next node
            list1 = list1.Next
        }else{
            // if list2 current node value smaller, do the same to list2
            tail.Next = list2
            list2 = list2.Next
        }
        // after linked the smaller node to the dummy list, move dummy node to the next node
        tail = tail.Next
    }
    // previously we have done with the both list are not empty
    // Now let's check if one of them are empty
    // then link the other one to the dummy listnode
    if list1 == nil{
        tail.Next = list2
    }else{
        tail.Next = list1
    }
    return dummy.Next
}
// @lc code=end
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	if list1 == nil {
// 	   return list2
// 	}
// 	if list2 == nil {
// 	   return list1
// 	}

// 	if list1.Val < list2.Val {
// 	   list1.Next = mergeTwoLists(list1.Next, list2)
// 	   return list1
// 	}
// 	list2.Next = mergeTwoLists(list1, list2.Next)
// 	return list2
// }

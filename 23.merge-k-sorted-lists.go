/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (47.93%)
 * Likes:    14414
 * Dislikes: 539
 * Total Accepted:    1.4M
 * Total Submissions: 3M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * You are given an array of k linked-lists lists, each linked-list is sorted
 * in ascending order.
 * 
 * Merge all the linked-lists into one sorted linked-list and return it.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: lists = [[1,4,5],[1,3,4],[2,6]]
 * Output: [1,1,2,3,4,4,5,6]
 * Explanation: The linked-lists are:
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * merging them into one sorted list:
 * 1->1->2->3->4->4->5->6
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: lists = []
 * Output: []
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: lists = [[]]
 * Output: []
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] is sorted in ascending order.
 * The sum of lists[i].length will not exceed 10^4.
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
func mergeKLists(lists []*ListNode) *ListNode {
	
	// for exception
	if lists == nil || len(lists) ==0 {
		return nil
	}

	// keep merge the fist 2 list of lists to a single listNode
	for len(lists) > 1{
		list1 := lists[0]
		list2 := lists[1]

		// pop the fist two listNode which need to be merged
		lists = lists[2:]
		// merged two lists
		merged := mergeTwoLists(list1, list2)

		// append the merged listNode back the lists, until the list length == 1
		// means every list has been merged to lists[0]
		lists = append(lists, merged)
	}
	
    return lists[0]
}

// merge two listNode based on problem 21
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
//
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode{
// 	if list1 == nil {
// 		return list2
// 	}

// 	if list2 == nil {
// 		return list1
// 	}

// 	if list1.Val < list2.Val{
// 		list1.Next = mergeTwoLists(list1.Next, list2)
// 		return list1
// 	}
// 	list2.Next = mergeTwoLists(list1, list2.Next)
// 	return list2
// }
// @lc code=end


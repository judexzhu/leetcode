/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 *
 * https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (68.78%)
 * Likes:    8566
 * Dislikes: 150
 * Total Accepted:    959.1K
 * Total Submissions: 1.4M
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * Given the root of a binary search tree, and an integer k, return the k^th
 * smallest value (1-indexed) of all the values of the nodes in the tree.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root = [3,1,4,null,2], k = 1
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root = [5,3,6,2,4,null,null,1], k = 3
 * Output: 3
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in the tree is n.
 * 1 <= k <= n <= 10^4
 * 0 <= Node.val <= 10^4
 * 
 * 
 * 
 * Follow up: If the BST is modified often (i.e., we can do insert and delete
 * operations) and you need to find the kth smallest frequently, how would you
 * optimize?
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func kthSmallest(root *TreeNode, k int) int {

	// create a stack
    stack := []*TreeNode{}
	// cur pointer to root 
	cur := root
	
	for cur != nil || len(stack) != 0 {
		// keep going to left until nil
		// add every Treenode passed to the stack
		for cur != nil{
			stack = append(stack, cur)
			cur = cur.Left
		}
		// when move to the very end of node left
		// cur == nil , fall back to the cur root node
		// move cur to cur.root by pop it from the stack 
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// every time we poped one node, k - 1
		// that will be the currently smallest node 
		// after pop k time (if k ==0), we find the node for answer
		k--

		if k == 0{
			return node.Val
		}
		// otherwise, let's try the right children
		if node.Right != nil {
			cur = node.Right
		}
	}
	// return - 1 when k is out of range (should never happened)
	return -1
}
// @lc code=end


/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (62.69%)
 * Likes:    11197
 * Dislikes: 211
 * Total Accepted:    1.6M
 * Total Submissions: 2.5M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given the root of a binary tree, return the level order traversal of its
 * nodes' values. (i.e., from left to right, level by level).
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root = [3,9,20,null,null,15,7]
 * Output: [[3],[9,20],[15,7]]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root = [1]
 * Output: [[1]]
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: root = []
 * Output: []
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in the tree is in the range [0, 2000].
 * -1000 <= Node.val <= 1000
 * 
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
func levelOrder(root *TreeNode) [][]int {

	// BFS
	res := [][]int{}

	// for root == nil 
	if root == nil {
		return res
	}

	// create a queue for BFS
	// add root node to the queue to start BFS
	queue := []*TreeNode{root}

	// go until queue empty
	for len(queue) > 0 {
		// get current length of this level
		// create an array to record the nodes Val in this level
		n := len(queue)
		level := []int{}

		for i := 0; i < n; i++ {
			// pop the very left node from the queue, add it's val to the level array
			// until all the node in this level have been taken care
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			// after pop current node, add its children to the queue (prepare for next level)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// finfish current level, for loop goes to next level
		res = append(res, level)
	}

	return res
}
// @lc code=end


/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
 *
 * https://leetcode.com/problems/invert-binary-tree/description/
 *
 * algorithms
 * Easy (72.69%)
 * Likes:    10026
 * Dislikes: 138
 * Total Accepted:    1.3M
 * Total Submissions: 1.7M
 * Testcase Example:  '[4,2,7,1,3,6,9]'
 *
 * Given the root of a binary tree, invert the tree, and return its root.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root = [4,2,7,1,3,6,9]
 * Output: [4,7,2,9,6,3,1]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root = [2,1,3]
 * Output: [2,3,1]
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
 * The number of nodes in the tree is in the range [0, 100].
 * -100 <= Node.val <= 100
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
func invertTree(root *TreeNode) *TreeNode {

	// for exception
	if root == nil {
		return nil
	}

	// recursive
	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
    
}
// @lc code=end


/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 *
 * https://leetcode.com/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (38.14%)
 * Likes:    11953
 * Dislikes: 592
 * Total Accepted:    853.6K
 * Total Submissions: 2.2M
 * Testcase Example:  '[1,2,3]'
 *
 * A path in a binary tree is a sequence of nodes where each pair of adjacent
 * nodes in the sequence has an edge connecting them. A node can only appear in
 * the sequence at most once. Note that the path does not need to pass through
 * the root.
 * 
 * The path sum of a path is the sum of the node's values in the path.
 * 
 * Given the root of a binary tree, return the maximum path sum of any
 * non-empty path.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root = [1,2,3]
 * Output: 6
 * Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 =
 * 6.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root = [-10,9,20,null,null,15,7]
 * Output: 42
 * Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 +
 * 7 = 42.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in the tree is in the range [1, 3 * 10^4].
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
func maxPathSum(root *TreeNode) int {

	res := root.Val

	max := func(a int, b int) int {
		if a > b {
			return a 
		}
		return b
	}

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		// base case of traverse tree
		if root == nil {
			return 0
		}
		// calc leftMax and rightMax node
		left := dfs(root.Left)
		right := dfs(root.Right)

		// if leftMax or rightMax is negative
		// we except leftMax and rightmax from current path
		left = max(left, 0)
		right = max(right, 0)

		// update res if we take subpath as a result
		// -> we dont split from parent node
		res = max(res, root.Val+left+right)
		
		//return maximum of sub-path to parent
		return root.Val + max(left, right)
	}

	dfs(root)

	return res
    
}
// @lc code=end


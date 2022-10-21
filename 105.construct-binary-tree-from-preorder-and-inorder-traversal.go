/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (60.07%)
 * Likes:    11315
 * Dislikes: 318
 * Total Accepted:    870.8K
 * Total Submissions: 1.4M
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * Given two integer arrays preorder and inorder where preorder is the preorder
 * traversal of a binary tree and inorder is the inorder traversal of the same
 * tree, construct and return the binary tree.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 * Output: [3,9,20,null,null,15,7]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: preorder = [-1], inorder = [-1]
 * Output: [-1]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= preorder.length <= 3000
 * inorder.length == preorder.length
 * -3000 <= preorder[i], inorder[i] <= 3000
 * preorder and inorder consist of unique values.
 * Each value of inorder also appears in preorder.
 * preorder is guaranteed to be the preorder traversal of the tree.
 * inorder is guaranteed to be the inorder traversal of the tree.
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
 func buildTree(preorder []int, inorder []int) *TreeNode {
    // base case, if empty return nil
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    
	// the first element from the preorder is must be the root
    id := index(inorder, preorder[0])
    
	// use the root ip to split the preorder and inorder
	// inorder, everything left to id is in the left
	// so whole left tree is [:id], total numbers of the nodes is id -1 
	// right tree is [id+1:](remove the root)
	// preorder, now we know the left tress start from [1]
	// total id -1 number, so [1: id -1 + 1 + 1] = [1:id+1]
	// the others are the right tree [id+1:]
    return &TreeNode{
        Val: preorder[0],
        Left: buildTree(preorder[1:id+1], inorder[:id]),
        Right: buildTree(preorder[id+1:], inorder[id+1:]),
    }
    
    
}

func index(nums []int, target int) int {
    for i, num := range nums {
        if num == target {
            return i
        }
    }
    return -1
}
// @lc code=end


/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (54.11%)
 * Likes:    20331
 * Dislikes: 1077
 * Total Accepted:    1.8M
 * Total Submissions: 3.3M
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * You are given an integer array height of length n. There are n vertical
 * lines drawn such that the two endpoints of the i^th line are (i, 0) and (i,
 * height[i]).
 * 
 * Find two lines that together with the x-axis form a container, such that the
 * container contains the most water.
 * 
 * Return the maximum amount of water a container can store.
 * 
 * Notice that you may not slant the container.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: height = [1,8,6,2,5,4,8,3,7]
 * Output: 49
 * Explanation: The above vertical lines are represented by array
 * [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the
 * container can contain is 49.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: height = [1,1]
 * Output: 1
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * n == height.length
 * 2 <= n <= 10^5
 * 0 <= height[i] <= 10^4
 * 
 * 
 */

// @lc code=start
func maxArea(height []int) int {

	l, r := 0, len(height) - 1
	res, area := 0, 0

	// Two pointer
	// calulate the area between two pointer with the min() and distance
	// if [l] < [r], move l to right 
	// else move r to left
	// compare to get the max one
	// O(n)
	for l < r {
		area = (r - l) * min( height[l], height[r])
		res = max(res, area)

		if height[l] < height[r]{
			l++
		}else{
			r--
		}
	}
    return res
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
// @lc code=end


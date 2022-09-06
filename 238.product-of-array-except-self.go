/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 *
 * https://leetcode.com/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (64.43%)
 * Likes:    14368
 * Dislikes: 823
 * Total Accepted:    1.3M
 * Total Submissions: 2.1M
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an integer array nums, return an array answer such that answer[i] is
 * equal to the product of all the elements of nums except nums[i].
 * 
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 * 
 * You must write an algorithm that runs in O(n) time and without using the
 * division operation.
 * 
 * 
 * Example 1:
 * Input: nums = [1,2,3,4]
 * Output: [24,12,8,6]
 * Example 2:
 * Input: nums = [-1,1,0,-3,3]
 * Output: [0,0,9,0,0]
 * 
 * 
 * Constraints:
 * 
 * 
 * 2 <= nums.length <= 10^5
 * -30 <= nums[i] <= 30
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 * 
 * 
 * 
 * Follow up: Can you solve the problem in O(1) extra space complexity? (The
 * output array does not count as extra space for space complexity analysis.)
 * 
 */
package leetcode
// @lc code=start
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	prefix := 1
	for i:=0 ;i < n; i++ {
		res[i] = prefix
		prefix *= nums[i]
	}
	postfix := 1
	for i:=n - 1; i >= 0 ; i--{
		res[i] *= postfix
		postfix *= nums[i]
	} 
	return res
}
// @lc code=end

// Brute force will excced time limit, not acceptable, O(N2)
// func productExceptSelf(nums []int) []int {
// 	res := []int{}

// 	for i, _ := range nums{
// 		n := 1
// 		for j, w :=range nums{
			
// 			if i != j {
// 				n = n * w
// 			}
// 		}
// 		res = append(res, n)
// 	}
// 	return res
// }
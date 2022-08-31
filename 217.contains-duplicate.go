/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 *
 * https://leetcode.com/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (61.13%)
 * Likes:    6274
 * Dislikes: 1012
 * Total Accepted:    2M
 * Total Submissions: 3.2M
 * Testcase Example:  '[1,2,3,1]'
 *
 * Given an integer array nums, return true if any value appears at least twice
 * in the array, and return false if every element is distinct.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3,1]
 * Output: true
 * Example 2:
 * Input: nums = [1,2,3,4]
 * Output: false
 * Example 3:
 * Input: nums = [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */
package leetcode

// @lc code=start
func containsDuplicate(nums []int) bool {
    // hashmap (go has no hashset data structure)
	// Time: O(nlogn) 
	// Space: O(n)
	hashmap := make(map[int]bool, len(nums))

	for _, n := range nums {
		if _, ok := hashmap[n]; ok {
			return true
		}
		hashmap[n] = true
	}
	return false
}
// @lc code=end

// func containsDuplicate(nums []int) bool {
    // Brute force
	// Time: O(n2) 
	// Space: O(1)
// 	n := len(nums)
// 	for i := 0; i < n ; i ++ {
// 		for j := 0; j < n ; j ++ {
// 			if i != j && nums[i] == nums[j]{
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func containsDuplicate(nums []int) bool {
//     // Sorting
// 	// Time: O(nlogn) 
// 	// Space: O(1)
// 	sort.Ints(nums)

// 	for i := 0; i < len(nums) - 1; i ++ {
// 		if nums[i] == nums[i+1]{
// 			return true
// 		}
// 	}

// 	return false

// }
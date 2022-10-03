/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (38.37%)
 * Likes:    17809
 * Dislikes: 1067
 * Total Accepted:    1.7M
 * Total Submissions: 4.5M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * There is an integer array nums sorted in ascending order (with distinct
 * values).
 * 
 * Prior to being passed to your function, nums is possibly rotated at an
 * unknown pivot index k (1 <= k < nums.length) such that the resulting array
 * is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]
 * (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3
 * and become [4,5,6,7,0,1,2].
 * 
 * Given the array nums after the possible rotation and an integer target,
 * return the index of target if it is in nums, or -1 if it is not in nums.
 * 
 * You must write an algorithm with O(log n) runtime complexity.
 * 
 * 
 * Example 1:
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 * Example 2:
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 * Example 3:
 * Input: nums = [1], target = 0
 * Output: -1
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums.length <= 5000
 * -10^4 <= nums[i] <= 10^4
 * All values of nums are unique.
 * nums is an ascending array that is possibly rotated.
 * -10^4 <= target <= 10^4
 * 
 * 
 */

// @lc code=start
func search(nums []int, target int) int {

	l,r,mid := 0,len(nums)-1,0

	for l <= r {
		mid = (l+r)/2

		if target == nums[mid]{
			return mid
		}

		// left sorted portion
		if  nums[l] <= nums[mid]{
			// if target greater than the mid value or less than the left value
			// the entire left portion can be ignored
			if target > nums[mid] || target < nums[l]{
				l = mid + 1
			// otherwise do the binary search in left portion
			} else {
				r = mid - 1
			}
			// right sorted portion
		}else{
			// if targe less than the mid value or greate than the very right value
			// then entire right portion can be ignored
			if target < nums[mid] || target > nums[r]{
				r = mid - 1
			// other wide do binary in right portion
			} else {
				l = mid + 1
			}
		}
	}
    return -1 
}
// @lc code=end


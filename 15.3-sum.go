/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
package leetcode

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	// iterate throght sorted slice nums
	for i, v := range nums{
		// after sort, the first number will not count as duplicate
		// so when i > 0 and num[i] == num[i-1], it's duplicates, skip to next one (continue) 
		if i > 0 && v == nums[i-1] {
			continue
		}
		// after we find the fist num, we just use the left of the slice to do 2sum II, by using two pointer or hashmap
		// two point for 2sum

		t := 0 - v

		// set the left point 
		l := i + 1
		//set the right point 
		r := len(nums) - 1

		// move the two pointers 
		for l < r {
			// Greater than t, move r - 1
			// less , move l + 1 
			if nums[l] + nums[r] == t {
				result = append(result, []int{v, nums[l], nums[r]})

				// after find the first combine, keep moving left pointer if left same value to avoid duplicate in result.
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			} 
			
			if nums[l] + nums[r] < t {
				l++
			} else {
				r--
				

			}
		}
	}

	return result
}
// @lc code=end

// func threeSum(nums []int) [][]int {
// 	sort.Ints(nums)
// 	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
// 	for index = 1; index < length-1; index++ {
// 		start, end = 0, length-1
// 		if index > 1 && nums[index] == nums[index -1]{
// 			start = index - 1
// 		}
// 		for start < index && end > index {
// 			if start > 0 && nums[start] == nums[start -1]{
// 				start++
// 				continue
// 			}
// 			if end < length-1 && nums[end] == nums[end+1]{
// 				end --
// 				continue
// 			}
// 			addNum = nums[start] + nums[end] + nums[index]
// 			if addNum == 0 {
// 				result = append(result, []int{nums[start], nums[index], nums[end]})
// 				start ++
// 				end --
// 			} else if addNum > 0 {
// 				end --
// 			} else {
// 				start ++
// 			}
// 		}
// 	}
// 	return result
// }


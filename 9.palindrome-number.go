/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (52.71%)
 * Likes:    6847
 * Dislikes: 2259
 * Total Accepted:    2.4M
 * Total Submissions: 4.5M
 * Testcase Example:  '121'
 *
 * Given an integer x, return true if x is palindrome integer.
 *
 * An integer is a palindrome when it reads the same backward as forward.
 *
 *
 * For example, 121 is a palindrome while 123 is not.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: x = 121
 * Output: true
 * Explanation: 121 reads as 121 from left to right and from right to left.
 *
 *
 * Example 2:
 *
 *
 * Input: x = -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: x = 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= x <= 2^31 - 1
 *
 *
 *
 * Follow up: Could you solve it without converting the integer to a string?
 */
package leetcode

import "strconv"

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	s := strconv.Itoa(x)
	length := len(s)

	for i := 0; i <= length/2; i++ {
		if s[i] != s[length -1 -i] {
			return false
		}
	}
	return true
}
// func isPalindrome(x int) bool {
// 	s := strconv.Itoa(x)
// 	r := []rune(s)
// 	for i, j := 0 , len(r)-1 ; i < j; i, j = i+1 , j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	if s == string(r) {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// @lc code=end


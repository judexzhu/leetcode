/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (37.03%)
 * Likes:    4923
 * Dislikes: 3903
 * Total Accepted:    1.4M
 * Total Submissions: 3.8M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 * 
 * Given two strings needle and haystack, return the index of the first
 * occurrence of needle in haystack, or -1 if needle is not part of haystack.
 * 
 * Clarification:
 * 
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 * 
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= haystack.length, needle.length <= 10^4
 * haystack and needle consist of only lowercase English characters.
 * 
 * 
 */
package leetcode
// @lc code=start
func strStr(haystack string, needle string) int {
	n := len(needle)

    if haystack == needle || n == 0 {
		return 0
	}

	for i := 0 ; i < len(haystack) - n + 1 ; i++ {
		if haystack[i:i + n] == needle {
			return i
		}
	}
	return -1 

}
// @lc code=end


/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (40.34%)
 * Likes:    9851
 * Dislikes: 3308
 * Total Accepted:    1.8M
 * Total Submissions: 4.5M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 *
 * Example 1:
 *
 *
 * Input: strs = ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: strs = ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] consists of only lowercase English letters.
 *
 *
 */
package leetcode

import "sort"

// @lc code=start
func longestCommonPrefix(strs []string) string {
	// First, sort the string slice
	sort.Strings(strs)

	// After sort. we only need compare the first word with the last word
	first_word := strs[0]	
	last_word := strs[len(strs)-1]
	prefix := ""

	for i:=0; i<len(first_word); i++ {
		// compare start from the first char of first word and last word, if same, try next char. once diff, break the for loop.
		if first_word[i] == last_word[i]{
			prefix = prefix + string(first_word[i])
		} else {
			break
		}
	}
    return prefix
}
// @lc code=end


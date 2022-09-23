/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 *
 * https://leetcode.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (58.58%)
 * Likes:    6653
 * Dislikes: 227
 * Total Accepted:    1.2M
 * Total Submissions: 2.1M
 * Testcase Example:  '"leetcode"'
 *
 * Given a string s, find the first non-repeating character in it and return
 * its index. If it does not exist, return -1.
 * 
 * 
 * Example 1:
 * Input: s = "leetcode"
 * Output: 0
 * Example 2:
 * Input: s = "loveleetcode"
 * Output: 2
 * Example 3:
 * Input: s = "aabb"
 * Output: -1
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= s.length <= 10^5
 * s consists of only lowercase English letters.
 * 
 * 
 */
package leetcode
// @lc code=start
func firstUniqChar(s string) int {
    arr := [26]int{}

	for _, v := range s {
		arr[v-'a'] ++
	}
	for i, _ := range s {
		if arr[s[i] - 'a'] == 1 {
			return i
		}
	}
	return -1
}
// @lc code=end


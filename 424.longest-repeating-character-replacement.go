/*
 * @lc app=leetcode id=424 lang=golang
 *
 * [424] Longest Repeating Character Replacement
 *
 * https://leetcode.com/problems/longest-repeating-character-replacement/description/
 *
 * algorithms
 * Medium (51.24%)
 * Likes:    6250
 * Dislikes: 245
 * Total Accepted:    334.5K
 * Total Submissions: 651.6K
 * Testcase Example:  '"ABAB"\n2'
 *
 * You are given a string s and an integer k. You can choose any character of
 * the string and change it to any other uppercase English character. You can
 * perform this operation at most k times.
 * 
 * Return the length of the longest substring containing the same letter you
 * can get after performing the above operations.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "ABAB", k = 2
 * Output: 4
 * Explanation: Replace the two 'A's with two 'B's or vice versa.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "AABABBA", k = 1
 * Output: 4
 * Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
 * The substring "BBBB" has the longest repeating letters, which is 4.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= s.length <= 10^5
 * s consists of only uppercase English letters.
 * 0 <= k <= s.length
 * 
 * 
 */

// @lc code=start
func characterReplacement(s string, k int) int {
    // Create an array to represent 26 char from 'A'~'Z'
	arr := [26]int{}
	res := 0
	// Using Slide windows
	// l, r start from the very left
	l,r := 0,0

	for r < len(s){
		// move r to right and count the r char to array
		arr[s[r] - 'A']++
		r++
		// when the max replaceable char count bigger than the k
		// move left to right 
		// also clean the removed l char count from array
		for r-l-Max(arr[:]) > k {
			arr[s[l] - 'A' ]--
			l++
		}
		// get the max res
		if r - l > res{
			res = r-l
		}
	}
	return res

}
// find the max count from currently array
func Max(arr []int) int {
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
// @lc code=end


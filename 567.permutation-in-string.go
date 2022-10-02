/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 *
 * https://leetcode.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (44.07%)
 * Likes:    7249
 * Dislikes: 239
 * Total Accepted:    486.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * Given two strings s1 and s2, return true if s2 contains a permutation of s1,
 * or false otherwise.
 * 
 * In other words, return true if one of s1's permutations is the substring of
 * s2.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s1 = "ab", s2 = "eidbaooo"
 * Output: true
 * Explanation: s2 contains one permutation of s1 ("ba").
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s1 = "ab", s2 = "eidboaoo"
 * Output: false
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= s1.length, s2.length <= 10^4
 * s1 and s2 consist of lowercase English letters.
 * 
 * 
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {

	// create a 26int arr to represent a-z
	// also can do with hashmap
	arr1 := [26]int{}
	// arr1 respent the count of all the s1 chars
	for _,v := range s1 {
		arr1[v - 'a']++
	}
	// makeing a slide window length same as s1
	l,r := 0,len(s1)-1
	// slide the windows, move to see if the arr2 matches s1
	// if match, s2 has s1 permutation, return true
	// if no match return false 
	for r < len(s2){
		arr2 := [26]int{}

		for _,v := range s2[l:r+1]{
			arr2[v - 'a']++
		}

		if arr1 == arr2 {
			return true
		}

		l++
		r++
	}
	
    
	return false
}
// @lc code=end


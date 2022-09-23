/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (42.96%)
 * Likes:    4889
 * Dislikes: 5997
 * Total Accepted:    1.6M
 * Total Submissions: 3.6M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * A phrase is a palindrome if, after converting all uppercase letters into
 * lowercase letters and removing all non-alphanumeric characters, it reads the
 * same forward and backward. Alphanumeric characters include letters and
 * numbers.
 * 
 * Given a string s, return true if it is a palindrome, or false otherwise.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "A man, a plan, a canal: Panama"
 * Output: true
 * Explanation: "amanaplanacanalpanama" is a palindrome.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "race a car"
 * Output: false
 * Explanation: "raceacar" is not a palindrome.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: s = " "
 * Output: true
 * Explanation: s is an empty string "" after removing non-alphanumeric
 * characters.
 * Since an empty string reads the same forward and backward, it is a
 * palindrome.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= s.length <= 2 * 10^5
 * s consists only of printable ASCII characters.
 * 
 * 
 */
 package leetcode

 import (
	"strings"
 )

// @lc code=start
func isPalindrome(s string) bool {
	if len(s)==0 {
		return true
	}

	s = cleanStrings(strings.ToLower(s))

	for i :=0; i < len(s)/2; i++ {
		if s[i] != s [len(s) -1 -i ]{
			return false
		}
	}
	return true
}

func cleanStrings(s string) string {
	var result strings.Builder
	for _,v := range s {
		if ('a' <= v && v <= 'z') || ('0' <= v && v <= '9'){
			result.WriteString(string(v))
		}
	}
	return result.String()
}
// @lc code=end


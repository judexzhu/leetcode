/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 *
 * https://leetcode.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (62.51%)
 * Likes:    6789
 * Dislikes: 238
 * Total Accepted:    1.6M
 * Total Submissions: 2.6M
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t, return true if t is an anagram of s, and false
 * otherwise.
 *
 * An Anagram is a word or phrase formed by rearranging the letters of a
 * different word or phrase, typically using all the original letters exactly
 * once.
 *
 *
 * Example 1:
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 * Example 2:
 * Input: s = "rat", t = "car"
 * Output: false
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length, t.length <= 5 * 10^4
 * s and t consist of lowercase English letters.
 *
 *
 *
 * Follow up: What if the inputs contain Unicode characters? How would you
 * adapt your solution to such a case?
 *
 */
package leetcode

// @lc code=start
func isAnagram(s string, t string) bool {

	hashmap := make(map[rune]int, len(s))

	if len(s) != len(t){
		return false
	}

	for _, i := range s {
		hashmap[i] += 1
	}

	for _, j := range t {
		hashmap[j] -= 1
		}

	for _, v := range hashmap {
		if v != 0 {
			return false
		}
	}
	return true
}
	

// @lc code=end

// func isAnagram(s string, t string) bool {

// 	w1 := strings.Split(s, "")
// 	w2 := strings.Split(t, "")

// 	sort.Strings(w1)
// 	sort.Strings(w2)

// 	if strings.Join(w1, "") == strings.Join(w2, ""){
// 		return true
// 	} 
//     return false
// }

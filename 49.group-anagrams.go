/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (65.37%)
 * Likes:    11407
 * Dislikes: 360
 * Total Accepted:    1.6M
 * Total Submissions: 2.4M
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings strs, group the anagrams together. You can return
 * the answer in any order.
 * 
 * An Anagram is a word or phrase formed by rearranging the letters of a
 * different word or phrase, typically using all the original letters exactly
 * once.
 * 
 * 
 * Example 1:
 * Input: strs = ["eat","tea","tan","ate","nat","bat"]
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 * Example 2:
 * Input: strs = [""]
 * Output: [[""]]
 * Example 3:
 * Input: strs = ["a"]
 * Output: [["a"]]
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= strs.length <= 10^4
 * 0 <= strs[i].length <= 100
 * strs[i] consists of lowercase English letters.
 * 
 * 
 */
package leetcode
// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// time： O(m*n) m为 slice的长度 ，n 为 string的平均长度
	// space： O(26*m)
	// 创建一个由 string slice 组成的 slice 来存放结果
    res := [][]string{}
	//创建一个 hashmap， key 为 代表 26个字母的slice
	// value为 string slice
	tmp := map[[26]int][]string{}

	//
	for _, s := range strs {
		chars := [26]int{}
		// 遍历每个string中的每个字符，并且根据26位字母数序
		//每包含一个字母就++
		// abc >>> [1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
		for _,c := range s {
			chars[c -'a']++
		}
		//根据字母顺序产生的slice，将 anagrams words加入到对应的 string slice中
		tmp[chars] = append(tmp[chars], s)
	}
	//将所有的结果 hashmap 中的 string的 slice，合并成一个 slice
	for _, v := range tmp {
		res = append(res, v)
	}
	return res
}
// @lc code=end


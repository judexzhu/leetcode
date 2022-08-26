/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (40.80%)
 * Likes:    15145
 * Dislikes: 745
 * Total Accepted:    2.6M
 * Total Submissions: 6.3M
 * Testcase Example:  '"()"'
 *
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 * 
 * An input string is valid if:
 * 
 * 
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "()"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "()[]{}"
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: s = "(]"
 * Output: false
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= s.length <= 10^4
 * s consists of parentheses only '()[]{}'.
 * 
 * 
 */

package leetcode
// @lc code=start
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	p := map[string]string {
		")": "(",
		"]": "[",
		"}": "{",
	}
	var stack []string

	 for _,c := range s {
		w := string(c)
		// 遇到左括号入栈，遇到右括号如果 栈不为空 并且栈尾为对应的左括号， 移除栈尾
		/*
		 If left bracket, push onto the stack, 
		 If right bracket is encountered, if the stack is not empty and the last in stack is the corresponding left bracket, 
		 remove the last in stack
		*/
		if v, ok := p[w] ; ok {
			if len(stack) > 0 && stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1]
			} else {
				// 如果栈为空，但是轮训到右括号，返回false
				// If the stack is empty, but the right parenthesis is encountered, return false
				return false
			}
		} else {
			//左括号入栈
			//left bracket, push onto the stack,
			stack = append(stack, w)
		}
		
	}
	//如果最后栈为空，说明所有的括号都按照顺序配对，返回 true，否则 false
	//If in the end, the stack is empty, indicating that all parentheses are matched in order, return true, otherwise return false
	return len(stack) ==0 
}

// @lc code=end


// func isValid(s string) bool {
// 	p := map[string]string {
// 		")": "(",
// 		"]": "[",
// 		"}": "{",
// 	}

// 	var stack []string

// 	 for _,c := range s {
// 		w := string(c)

// 		if v, ok := p[w] ; ok {
// 			if len(stack) > 0 && stack[len(stack)-1] == v {
// 				stack = stack[:len(stack)-1]
// 			} else {
// 				return false
// 			}
// 		} else {
// 			stack = append(stack, w)
// 		}
		
// 	}
// 	if len(stack) ==0 {
// 		return true
// 	} else {
// 		return false
// 	}

// }
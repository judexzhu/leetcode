/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 *
 * https://leetcode.com/problems/valid-sudoku/description/
 *
 * algorithms
 * Medium (56.27%)
 * Likes:    6289
 * Dislikes: 789
 * Total Accepted:    857.2K
 * Total Submissions: 1.5M
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be
 * validated according to the following rules:
 *
 *
 * Each row must contain the digits 1-9 without repetition.
 * Each column must contain the digits 1-9 without repetition.
 * Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9
 * without repetition.
 *
 *
 * Note:
 *
 *
 * A Sudoku board (partially filled) could be valid but is not necessarily
 * solvable.
 * Only the filled cells need to be validated according to the mentioned
 * rules.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: board =
 * [["5","3",".",".","7",".",".",".","."]
 * ,["6",".",".","1","9","5",".",".","."]
 * ,[".","9","8",".",".",".",".","6","."]
 * ,["8",".",".",".","6",".",".",".","3"]
 * ,["4",".",".","8",".","3",".",".","1"]
 * ,["7",".",".",".","2",".",".",".","6"]
 * ,[".","6",".",".",".",".","2","8","."]
 * ,[".",".",".","4","1","9",".",".","5"]
 * ,[".",".",".",".","8",".",".","7","9"]]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: board =
 * [["8","3",".",".","7",".",".",".","."]
 * ,["6",".",".","1","9","5",".",".","."]
 * ,[".","9","8",".",".",".",".","6","."]
 * ,["8",".",".",".","6",".",".",".","3"]
 * ,["4",".",".","8",".","3",".",".","1"]
 * ,["7",".",".",".","2",".",".",".","6"]
 * ,[".","6",".",".",".",".","2","8","."]
 * ,[".",".",".","4","1","9",".",".","5"]
 * ,[".",".",".",".","8",".",".","7","9"]]
 * Output: false
 * Explanation: Same as Example 1, except with the 5 in the top left corner
 * being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it
 * is invalid.
 *
 *
 *
 * Constraints:
 *
 *
 * board.length == 9
 * board[i].length == 9
 * board[i][j] is a digit 1-9 or '.'.
 *
 *
 */
package leetcode

import "fmt"

// @lc code=start
func isValidSudoku(board [][]byte) bool {
    for row := 0; row < 9; row++ {
		if !isValidSudokuRow(board,row){
			return false
		}
	}

    for col := 0; col < 9; col++ {
		if !isValidSudokuCol(board,col){
			return false
		}
	}

    for square := 0; square < 9; square++ {
		if !isValidSudokuSquare(board,square){
			return false
		}
	}
	return true
}

func isValidSudokuRow(board [][]byte, row int) bool {
	nums := [10]bool{}
	for col := 0; col < 9; col++{
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			fmt.Println("Invalid row: ", row)
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidSudokuCol(board [][]byte, col int) bool {
	nums := [10]bool{}
	for row := 0; row < 9; row++{
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			fmt.Println("Invalid col: ", col)
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidSudokuSquare(board [][]byte, square int) bool{
	nums := [10]bool{}
	row := (square / 3) * 3
	col := (square % 3) * 3

	for srow :=0; srow < 3; srow++{
		for scol := 0; scol < 3; scol++{
			n := convertToNumber(board[row + srow][col + scol])
			if n < 0 {
				continue
			}
			if nums[n] {
				fmt.Println("Invalid Square: ", square)
				fmt.Printf("Found duplicate number %d on row %d, col %d", n, row+srow, col+scol)
				return false
			}
			nums[n] = true
		}
	}
	return true
}


func convertToNumber(b byte) int {
	if b == '.'{
		return -1 
	}
	return int(b -'0')
}

// @lc code=end


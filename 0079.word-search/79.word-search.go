package main

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	letters := []byte(word)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			letter := letters[0]
			if letter == board[i][j] {
				if search(board, letters, used, m, n, 1, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func search(board [][]byte, letters []byte, used [][]bool, m, n, cur, i, j int) bool {
	if cur == len(letters) {
		return true
	}
	used[i][j] = true
	if i > 0 && !used[i-1][j] && letters[cur] == board[i-1][j] {
		if search(board, letters, used, m, n, cur+1, i-1, j) {
			return true
		}
	}
	if j > 0 && !used[i][j-1] && letters[cur] == board[i][j-1] {
		if search(board, letters, used, m, n, cur+1, i, j-1) {
			return true
		}
	}
	if i < m-1 && !used[i+1][j] && letters[cur] == board[i+1][j] {
		if search(board, letters, used, m, n, cur+1, i+1, j) {
			return true
		}
	}
	if j < n-1 && !used[i][j+1] && letters[cur] == board[i][j+1] {
		if search(board, letters, used, m, n, cur+1, i, j+1) {
			return true
		}
	}
	used[i][j] = false
	return false
}

// @lc code=end

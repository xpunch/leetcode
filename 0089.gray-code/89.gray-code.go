package main

/*
 * @lc app=leetcode id=89 lang=golang
 *
 * [89] Gray Code
 */

// @lc code=start
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	total := 1 << n
	return dfs(total, n, true, 0, make([]int, 0, total))
}

func dfs(total int, len int, step bool, n int, result []int) []int {
	if total == 0 {
		return result
	}
	result = append(result, n)

	if step {
		n = n ^ 1
	} else {
		i := 0
		for ; i <= len-1; i++ {
			tmp := n | (1 << i)
			if 0 == n^tmp {
				break
			}
		}
		if i == len-1 {
			n = n ^ 1
		} else {
			n = n ^ (1 << (i + 1))
		}
	}

	return dfs(total-1, len, !step, n, result)
}

// @lc code=end

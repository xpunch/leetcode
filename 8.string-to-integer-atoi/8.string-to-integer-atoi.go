/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
func myAtoi(str string) int {
	max := 2<<30 - 1
	min := -(2 << 30)
	var result, start, end int
	var signFound, startFound, endFound bool
	sign := 1
	for i := 0; i < len(str); i++ {
		if !startFound {
			if signFound {
				if str[i] == '0' {
					continue
				} else if str[i] > '0' && str[i] <= '9' {
					start = i
					startFound = true
				} else {
					break
				}
			}
			if str[i] == ' ' {
				continue
			}
			if str[i] == '0' {
				signFound = true
				continue
			}
			if str[i] == '-' || str[i] == '+' {
				if str[i] == '-' {
					sign = -1
				}
				signFound = true
				continue
			} else if str[i] >= '0' && str[i] <= '9' {
				start = i
				end = i
				startFound = true
				endFound = true
			} else {
				break
			}
		} else {
			if str[i] >= 48 && str[i] <= 57 {
				end = i
				endFound = true
			} else {
				break
			}
		}
	}
	if !startFound || !endFound {
		return 0
	}
	if end-start > 10 {
		if sign == 1 {
			return max
		} else {
			return min
		}
	}
	j := 1
	for i := 0; i <= end-start; i++ {
		result += int(str[end-i]-48) * j * sign
		j = 10 * j
		if result < min {
			return min
		}
		if result > max {
			return max
		}
	}
	return result
}

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
func romanToInt(s string) int {
	maps := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var value int
	for i := 0; i < len(s); i++ {
		c := s[i]
		v, ok := maps[c]
		if !ok {
			return 0
		}
		if i != len(s)-1 {
			big := s[i+1]
			switch c {
			case 'I':
				if big == 'V' {
					v = 4
					i++
				} else if big == 'X' {
					v = 9
					i++
				}
				break
			case 'X':
				if big == 'L' {
					i++
					v = 40
				} else if big == 'C' {
					v = 90
					i++
				}
				break
			case 'C':
				if big == 'D' {
					v = 400
					i++
				} else if big == 'M' {
					v = 900
					i++
				}
				break
			}
		}
		value += v
	}
	return value
}

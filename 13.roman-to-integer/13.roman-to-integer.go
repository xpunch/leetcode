/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
func romanToInt(s string) int {
	var value int
	for i := 0; i < len(s); i++ {
		var v int
		switch s[i] {
		case 'I':
			v = 1
		case 'V':
			v = 5
		case 'X':
			v = 10
		case 'L':
			v = 50
		case 'C':
			v = 100
		case 'D':
			v = 500
		case 'M':
			v = 1000
		default:
			return 0
		}
		if i != len(s)-1 {
			switch s[i : i+2] {
			case "IV":
				v = 4
				i++
				break
			case "IX":
				v = 9
				i++
				break
			case "XL":
				v = 40
				i++
				break
			case "XC":
				v = 90
				i++
				break
			case "CD":
				v = 400
				i++
			case "CM":
				v = 900
				i++
			}
		}
		value += v
	}
	return value
}

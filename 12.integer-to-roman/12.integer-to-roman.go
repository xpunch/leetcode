package main

/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */
func intToRoman(num int) string {
	var roman []byte
	for i := 1; ; i *= 10 {
		var one, five, ten byte
		if i == 1 {
			one = 'I'
			five = 'V'
			ten = 'X'
		} else if i == 10 {
			one = 'X'
			five = 'L'
			ten = 'C'
		} else if i == 100 {
			one = 'C'
			five = 'D'
			ten = 'M'
		} else if i == 1000 {
			one = 'M'
		}

		var cur []byte
		j := num % 10

		switch j {
		case 1:
			cur = []byte{one}
		case 2:
			cur = []byte{one, one}
		case 3:
			cur = []byte{one, one, one}
		case 4:
			cur = []byte{one, five}
		case 5:
			cur = []byte{five}
		case 6:
			cur = []byte{five, one}
		case 7:
			cur = []byte{five, one, one}
		case 8:
			cur = []byte{five, one, one, one}
		case 9:
			cur = []byte{one, ten}
		}
		roman = append(cur, roman...)
		if num/10 == 0 {
			break
		}
		num = num / 10
	}
	return string(roman)
}

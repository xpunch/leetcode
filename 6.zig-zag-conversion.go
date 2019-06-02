package main

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	if n <= numRows {
		return s
	}
	zLen := 2*numRows - 2
	zCount := (n-1)/zLen + 1
	lastZLen := n % zLen
	rowLens := make([]int, numRows)
	for i := 0; i < numRows; i++ {
		if i == 0 {
			rowLens[i] = zCount
		} else if i == numRows-1 {
			rowLens[i] = zCount - 1
			if lastZLen == 0 || lastZLen > i {
				rowLens[i] = rowLens[i] + 1
			}
		} else {
			rowLens[i] = (zCount - 1) * 2
			if lastZLen == 0 || lastZLen > 2*numRows-2-i {
				rowLens[i] = rowLens[i] + 2
			} else if lastZLen <= 2*numRows-2-i && lastZLen > i {
				rowLens[i] = rowLens[i] + 1
			}
		}
	}
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		if i < zCount {
			result[i] = s[i*zLen]
		} else if i >= n-rowLens[numRows-1] {
			result[i] = s[(i-n+rowLens[numRows-1])*zLen+numRows-1]
		} else {
			r := 1
			last := rowLens[0]
			for ; r < numRows-1; r++ {
				if i < last+rowLens[r] {
					break
				}
				last = last + rowLens[r]
			}
			li := (i - last) / 2
			if (i-last)%2 == 0 {
				result[i] = s[li*zLen+r]
			} else {
				result[i] = s[li*zLen+zLen-r]
			}
		}
	}
	return string(result)
}

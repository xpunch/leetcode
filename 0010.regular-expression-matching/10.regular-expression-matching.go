package main

func split(p string) ([]string, bool) {
	segments := make([]string, 0)
	for i, j := 0, 0; i < len(p); i++ {
		if p[i] == '*' {
			if j == i {
				return nil, false
			}
			if i > j+1 {
				segments = append(segments, p[j:i-1])
			}
			segments = append(segments, p[i-1:i+1])
			j = i + 1
		}
		if i == len(p)-1 && j <= i {
			segments = append(segments, p[j:i+1])
		}
	}
	return segments, true
}

func match(s string, p string) bool {
	for i := 0; i < len(s); i++ {
		if p[i] != '.' && s[i] != p[i] {
			return false
		}
	}
	return true
}

func min(i int, j int) int {
	if i > j {
		return j
	}
	return i
}

/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	pSegments, valid := split(p)
	if !valid {
		return false
	}
	pSegCount := len(pSegments)
	sSegments := make([]int, pSegCount)
	for i := 0; i < pSegCount; i++ {
		sSegments[i] = -1
	}
	var si, pi, tmpIndex, tmpMatchedSegIndex int
	var matched, rollback bool
	pi = -1
	si = -1
	for !matched {
		tmpMatchedSegIndex = pi + 1
		tmpIndex = si + 1
		for i := tmpMatchedSegIndex; i < pSegCount; i++ {
			if pSegments[i][len(pSegments[i])-1] != '*' {
				if rollback {
					if i == 0 || pi == i-1 {
						return false
					}
					tmpIndex -= sSegments[i-1]
					i -= 2
					continue
				}
				sSegments[i] = len(pSegments[i])
				if tmpIndex+sSegments[i] <= len(s) && match(s[tmpIndex:tmpIndex+sSegments[i]], pSegments[i]) {
					tmpMatchedSegIndex = i
					tmpIndex += sSegments[i]
					if i == pSegCount-1 {
						if tmpIndex == len(s) {
							return true
						} else if pi == i-1 {
							return false
						}
					}
					if i == 0 || pi == i-1 {
						pi = i
						si = tmpIndex - 1
					}
					if i == pSegCount-1 {
						tmpIndex -= sSegments[i]
						tmpIndex -= sSegments[i-1]
						i -= 2
						rollback = true
						continue
					}
				} else {
					if i == pSegCount-1 && tmpIndex+sSegments[i] == len(s) {
						return false
					}
					if i == 0 || pi == i-1 {
						return false
					}
					tmpIndex -= sSegments[i-1]
					i -= 2
					rollback = true
					continue
				}
			} else {
				if rollback {
					rollback = false
					if sSegments[i] != 0 {
						if tmpIndex < 0 && sSegments[i] <= 1 {
							return false
						}
						sSegments[i]--
						tmpMatchedSegIndex = i
						tmpIndex += sSegments[i]
						if sSegments[i] == 0 && pi == i-1 {
							pi = i
							si = tmpIndex - 1
						}
					} else {
						if i == 0 || pi == i-1 {
							return false
						}
						rollback = true
						tmpIndex -= sSegments[i-1]
						i -= 2
					}
					continue
				} else {
					sSegments[i] = 0
				}
				sp := pSegments[i][0]
				for tmpIndex+sSegments[i] < len(s) && (sp == '.' || sp == s[tmpIndex+sSegments[i]]) {
					sSegments[i]++
				}
				tmpIndex += sSegments[i]
				tmpMatchedSegIndex = i
				if i == pSegCount-1 {
					if i != 0 && tmpIndex < len(s) {
						rollback = true
						tmpIndex -= sSegments[i]
						tmpIndex -= sSegments[i-1]
						i -= 2
						continue
					}
					if tmpIndex == len(s) {
						return true
					} else if pi == i-1 {
						return false
					}
				}
				if i == pi+1 && sSegments[i] == 0 {
					pi = i
				}
			}
			rollback = false
		}
	}
	return true
}

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
	// fmt.Printf("s: %s, p: %s %v\n", s, p, pSegments)
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
					// fmt.Printf("[X]s[%d]: %s,  p[%d]: %s %v\n", tmpIndex, s[0:tmpIndex], i-1, pSegments[i-1], sSegments)
					// tmpIndex -= sSegments[i]
					tmpIndex -= sSegments[i-1]
					if tmpIndex < 0 {
						tmpIndex = 0
					}
					i -= 2
					continue
				}
				sSegments[i] = len(pSegments[i])
				if tmpIndex+sSegments[i] <= len(s) && match(s[tmpIndex:tmpIndex+sSegments[i]], pSegments[i]) {
					// fmt.Printf("[A]s[%d]: %s,  p[%d]: %s, %v\n", tmpIndex, s[0:tmpIndex+sSegments[i]], i, pSegments[i], sSegments)
					tmpMatchedSegIndex = i
					tmpIndex += sSegments[i]
					if i == pSegCount-1 {
						if tmpIndex == len(s) {
							// fmt.Printf("[S1]s[%d]: %s,  p[%d]: %s\n", tmpIndex, s[0:tmpIndex], i, pSegments[i])
							return true
						} else if pi == i-1 {
							// fmt.Printf("[N0]s[%d]: %s,  p[%d]: %s\n", tmpIndex, s[0:tmpIndex], i, pSegments[i])
							return false
						}
					}
					// ?
					if i == 0 || pi == i-1 {
						pi = i
						si = tmpIndex - 1
						// fmt.Printf("[M1]s[%d]: %s,  p[%d]: %s\n", tmpIndex, s[0:tmpIndex], i, pSegments[i])
					}
					if i == pSegCount-1 {
						tmpIndex -= sSegments[i]
						tmpIndex -= sSegments[i-1]
						if tmpIndex < 0 {
							tmpIndex = 0
						}
						i -= 2
						rollback = true
						continue
					}
				} else {
					if i == pSegCount-1 && tmpIndex+sSegments[i] == len(s) {
						return false
					}
					// dest := min(len(s), tmpIndex+len(pSegments[i]))
					if i == 0 || pi == i-1 {
						// fmt.Printf("[N1]s[%d]: %s,  p[%d]: %s\n", tmpIndex, s[0:dest], i, pSegments[i])
						return false
					}
					// fmt.Printf("[R]s[%d]: %s,  p[%d]: %s %v\n", tmpIndex, s[0:tmpIndex], i-1, pSegments[i-1], sSegments)
					// tmpIndex -= sSegments[i]
					tmpIndex -= sSegments[i-1]
					if tmpIndex < 0 {
						tmpIndex = 0
					}
					i -= 2
					rollback = true
					continue
				}
			} else {
				if rollback {
					rollback = false
					if sSegments[i] != 0 {
						if tmpIndex < 0 && sSegments[i] <= 1 {
							// fmt.Printf("[E]s[%d] %d, %v\n", tmpIndex, i, sSegments)
							return false
						}
						sSegments[i]--
						tmpMatchedSegIndex = i
						tmpIndex += sSegments[i]
						if sSegments[i] == 0 && pi == i-1 {
							pi = i
							si = tmpIndex - 1
							// fmt.Printf("[M3]s[%d]: %s,  p[%d]: %s\n", tmpIndex, s[0:tmpIndex], i, pSegments[i])
						} else if tmpIndex >= 0 {
							// fmt.Printf("[C1]s[%d]: %s,  p[%d]: %s, %v:%v\n", tmpIndex, s[0:tmpIndex], i, pSegments[i], rollback, sSegments)
						}
					} else {
						if i == 0 || pi == i-1 {
							return false
						}
						rollback = true
						tmpIndex -= sSegments[i-1]
						if tmpIndex < 0 {
							tmpIndex = 0
						}
						i -= 2
						if tmpIndex > 0 {
							// fmt.Printf("[r]s[%d]: %s,  p[%d]: %s, %v:%v\n", tmpIndex, s[0:tmpIndex], i+1, pSegments[i+1], rollback, sSegments)
						}
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
				// fmt.Printf("[B]s[%d]: %s,  p[%d]: %s, si: %d, pi: %d, %v\n", tmpIndex, s[0:tmpIndex], i, pSegments[i], si, pi, sSegments)
				if i == pSegCount-1 {
					if i != 0 && tmpIndex < len(s) {
						rollback = true
						// fmt.Printf("[Q]s[%d]: %s,  p[%d]: %s, %v:%v\n", tmpIndex, s[0:tmpIndex], i, pSegments[i], rollback, sSegments)
						tmpIndex -= sSegments[i]
						tmpIndex -= sSegments[i-1]
						if tmpIndex < 0 {
							tmpIndex = 0
						}
						i -= 2
						continue
					}
					if tmpIndex == len(s) {
						// fmt.Printf("[S2]s[%d]: %s,  p[%d]: %s\n", tmpIndex, s[0:tmpIndex], i, pSegments[i])
						return true
					} else if pi == i-1 {
						// fmt.Printf("[N4]s[%d]: %s,  p[%d]: %s\n", tmpIndex, s[0:tmpIndex], i, pSegments[i])
						return false
					}
				}
				if i == pi+1 && sSegments[i] == 0 {
					pi = i
					// fmt.Printf("[M2]s[%d]: %s,  p[%d]: %s\n", tmpIndex, s[0:tmpIndex], i, pSegments[i])
				}
			}
			rollback = false
		}
		// fmt.Printf("[T]s[%d]: %s, pi: %d, s: %v\n", si, s[0:si+1], pi, sSegments)
	}
	return true
}

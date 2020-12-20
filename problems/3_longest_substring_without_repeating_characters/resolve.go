package main

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	startInd := 0
	curMap := make(map[rune]int)
	for i, r := range s {
		if latest, ok := curMap[r]; ok && latest >= 0 {
			length := i - startInd
			if length > maxLen {
				maxLen = length
			}
			startInd = latest + 1
			for k, v := range curMap {
				if v <= latest {
					// logical delete
					curMap[k] = -1
				}
			}
			curMap[r] = i
			continue
		}
		curMap[r] = i
	}

	// last substring
	length := len(s) - startInd
	if length > maxLen {
		maxLen = length
	}
	return maxLen
}

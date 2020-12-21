package main

func longestPalindrome(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		newAns := getLongest(s, i)
		if len(newAns) > len(ans) {
			ans = newAns
		}
	}
	return ans
}

func getLongest(s string, start int) string {
	// odd
	l := start
	r := start
	for {
		if l < 0 || r >= len(s) {
			break
		}
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}
	oddMax := s[l+1 : r]

	// even
	l = start
	r = start + 1
	evenMax := ""
	for {
		if l < 0 || r >= len(s) {
			break
		}
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}

	if l+1 < r-1 {
		evenMax = s[l+1 : r]
	}

	if len(oddMax) > len(evenMax) {
		return oddMax
	}
	return evenMax
}

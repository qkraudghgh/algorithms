package backtrack

func anagram(s1 string, s2 string) bool {
	c1, c2 := make([]int, 26), make([]int, 26)
	for i := 0; i < len(s1); i++ {
		pos := rune(s1[i]) - rune('a')
		c1[pos] = c1[pos] + 1
	}

	for i := 0; i < len(s2); i++ {
		pos := rune(s2[i]) - rune('a')
		c2[pos] = c2[pos] + 1
	}

	j := 0
	stillOK := true

	for j < 26 && stillOK {
		if c1[j] == c2[j] {
			j += 1
		} else {
			stillOK = false
		}
	}

	return stillOK
}

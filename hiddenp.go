package piscine

func HiddeP(s1, s2 string) int {
	l := len(s1)
	count := 0
	for _, r := range s1 {
		for i, k := range s2 {
			if r == k {
				count++
				s2 = s2[i+1:]
				break
			}
		}
	}
	if count == l {
		return 1
	}

	return 0
}

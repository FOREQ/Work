package piscine

func RepeatAlpha(s string) string {
	if len(s) == 0 {
		return s
	}
	var s1 string
	n := 1

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			n = int(r - 'a')
			for i := 0; i < n; i++ {
				s1 = s1 + string(r)
			}
		}
		if r >= 'A' && r <= 'Z' {
			n = int(r - 'A')
			for i := 0; i < n; i++ {
				s1 = s1 + string(r)
			}
		}

		s1 = s1 + string(r)
	}
	return s1
}

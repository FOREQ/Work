package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(n int, base string) {
	if len(base) < 2 {
		return
	}

	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return
			}
		}
	}

	if n < 0 {
		z01.PrintRune('-')
		if n == -n {
			baseLen := len(base)
			quotient := -(n / baseLen)
			remainder := -(n % baseLen)
			if quotient != 0 {
				PrintNbrBase(quotient, base)
			}
			z01.PrintRune(rune(base[remainder]))
			return
		}
		n = -n
	}

	if n == 0 {
		z01.PrintRune(rune(base[n]))
	} else {
		PrintNbrBase(n/len(base), base)
		z01.PrintRune(rune(base[n%len(base)]))
	}
}

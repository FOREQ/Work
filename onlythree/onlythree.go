package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	out := os.Args[1]

	if !(len(args) == 1) {
		return
	}
	rout := []rune(out)

	fcount := 0
	scount := 0

	for i := 0; i < len(rout)-1; i++ {
		if rout[i] == ' ' {
			fcount++
		} else {
			break
		}
	}

	for j := len(rout) - 1; j > 0; j-- {
		if rout[j] == ' ' {
			scount++
		} else {
			break
		}
	}

	for k := fcount; k < len(rout)-scount; k++ {
		if rout[k] == ' ' {
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			for k+1 < len(rout)-scount && rout[k+1] == ' ' {
				k++
			}
		}
		z01.PrintRune(rout[k])
	}
	z01.PrintRune('\n')
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	sentence := []rune(arg[0])
	frune := rune(arg[1][0])
	srune := rune(arg[2][0])

	if len(arg) != 3 {
		return
	}

	for i, r := range sentence {
		if r == frune {
			sentence[i] = srune
		} else {
			continue
		}
	}

	for _, s := range sentence {
		z01.PrintRune(s)
	}
}

// func searchReplace()string {
// 	args:= os.Args
// 	word:= [1]args
// 	a:= [2]args
// 	b:= [3]args
// 	for i:=0; i<len(word); i++ {
// 		if string(word[])
// 	}
// }

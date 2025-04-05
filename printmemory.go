package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	for i, byt := range arr {
		printHex(byt)
		z01.PrintRune(' ')

		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		}
	}

	z01.PrintRune('\n')

	for _, byt := range arr {
		if byt < 32 || byt > 126 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(byt))
		}
	}
	z01.PrintRune('\n')
}

func printHex(n byte) {
	hexChars := "0123456789abcdef"
	z01.PrintRune(rune(hexChars[n/16]))
	z01.PrintRune(rune(hexChars[n%16]))
}

// import "github.com/01-edu/z01"

// func PrintMemory(arr [10]byte) {
// 	hex := "0123456789abcdef"
// 	for i := 0; i < len(arr); i++ {
// 		z01.PrintRune(rune(hex[arr[i]>>4]))
// 		z01.PrintRune(rune(hex[arr[i]&15]))
// 		z01.PrintRune(' ')
// 		if (i-3)%4 == 0 || i == len(arr)-1 {
// 			z01.PrintRune('\n')
// 		}
// 	}
// 	for _, b := range arr {
// 		if b < 32 || b > 128 {
// 			z01.PrintRune('.')
// 		} else {
// 			z01.PrintRune(rune(b))
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

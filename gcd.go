package piscine

func Gcd(a, b uint) uint {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
		// a, b = b, a%b
	}
	return a

	// if a > b {
	// 	for i := b; i > 0; i-- {
	// 		if a%i == 0 && b%i == 0 {
	// 			return i
	// 		}
	// 	}
	// }
	// if a < b {
	// 	for i := a; i > 0; i-- {
	// 		if b%i == 0 && a%i == 0 {
	// 			return i
	// 		}
	// 	}
	// }

	// return a
}

package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	len1 := len(slice1)
	len2 := len(slice2)
	maxLen := len1
	result := make([]int, 0, len1+len2)

	if len2 > maxLen {
		maxLen = len2
	}

	for i := maxLen - 1; i >= 0; i-- {
		if i < len1 {
			result = append(result, slice1[i])
		}
		if i < len2 {
			result = append(result, slice2[i])
		}
	}

	return result
}

// func RevConcatAlternate(slice1, slice2 []int) []int {
// 	slice3 := make([]int, 0)
// 	l1 := len(slice1)
// 	l2 := len(slice2)

// 	if l1 > l2 && l2 != 0 {
// 		temp1 := slice1[l2:]
// 		temp2 := slice1[:l2]
// 		for i := l1 - l2 - 1; i >= 0; i-- {
// 			slice3 = append(slice3, temp1[i])
// 		}
// 		for i := l2 - 1; i >= 0; i-- {
// 			slice3 = append(slice3, temp2[i])
// 			slice3 = append(slice3, slice2[i])
// 		}
// 	}

// 	if l1 < l2 && l1 != 0 {
// 		temp1 := slice2[l1:]
// 		temp2 := slice2[:l1]
// 		for i := l2 - l1 - 1; i >= 0; i-- {
// 			slice3 = append(slice3, temp1[i])
// 		}
// 		for i := l1 - 1; i >= 0; i-- {
// 			slice3 = append(slice3, slice1[i])
// 			slice3 = append(slice3, temp2[i])
// 		}
// 	}

// 	if l1 == l2 {
// 		for i := l1 - 1; i >= 0; i-- {
// 			slice3 = append(slice3, slice1[i])
// 			slice3 = append(slice3, slice2[i])
// 		}
// 	}

// 	if l1 == 0 {
// 		for i := l2 - 1; i >= 0; i-- {
// 			slice3 = append(slice3, slice2[i])
// 		}
// 	}

// 	if l2 == 0 {
// 		for i := l1 - 1; i >= 0; i-- {
// 			slice3 = append(slice3, slice1[i])
// 		}
// 	}

// 	return slice3
// }

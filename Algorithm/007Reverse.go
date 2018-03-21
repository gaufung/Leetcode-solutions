func reverse(x int) int {
	if x == 0 {
		return 0
	} else if x > 0 {
		return reversePositive(x)
	} else {
		return reverseNegative(x)
	}
}

func reversePositive(x int) int {
	digits := process(digitsOfPositiveInt(x))
	maxDigits := digitsReverse(digitsOfPositiveInt32())
	if len(digits) > len(maxDigits) {
		return 0
	} else if len(maxDigits) == len(digits) {
		result := 0
		isBigger := false
		for idx := 0; idx < len(maxDigits); idx++ {
			d1 := maxDigits[idx]
			d2 := digits[idx]
			if d2 > d1 {
				if !isBigger {
					return 0
				}else{
					result = result*10 + d2
				}
			} else {
				if d2 < d1 {
					isBigger= true
				}
				result = result*10 + d2
			}
		}
		return result
	} else {
		result := 0
		for _, digit := range digits {
			result = result*10 + digit
		}
		return result
	}

}

func reverseNegative(x int) int {
	digits := process(digitsOfNegativeInt(x))
	minDigits := digitsReverse(digitsOfNegativeInt32())
	if len(digits) > len(minDigits) {
		return 0
	} else if len(digits) == len(minDigits) {
		result := 0
		isBigger := false
		for idx := 0; idx < len(minDigits); idx++ {
			d1 := minDigits[idx]
			d2 := digits[idx]
			if d2 > d1 {
				if !isBigger {
					return 0
				} else {
					result = result*10 + d2
				}

			} else {
				if d2 < d1 {
					isBigger = true
				}
				result = result*10 + d2
			}
		}
		return -1 * result
	} else {
		result := 0
		for _, digit := range digits {
			result = result*10 + digit
		}
		return -1 * result
	}
}

func process(data []int) []int {
	idx := 0
	for _, v := range data {
		if v != 0 {
			break
		} else {
			idx++
		}
	}
	return data[idx:]
}

func digitsOfPositiveInt(x int) []int {
	result := make([]int, 0)
	for x > 0 {
		result = append(result, int(x%10))
		x = x / 10
	}
	return result
}

func digitsOfNegativeInt(x int) []int {
	result := make([]int, 0)
	for x < 0 {
		digit := x % 10
		result = append(result, -1*digit)
		x = x / 10
	}
	return result
}

func digitsReverse(data []int) []int {
	left, right := 0, len(data)-1
	for left < right {
		data[left], data[right] = data[right], data[left]
		left++
		right--
	}
	return data
}

func digitsOfPositiveInt32() []int {
	max, _ := maxMinInt32()
	result := make([]int, 0)
	for max > 0 {
		result = append(result, int(max%10))
		max = max / 10
	}
	return result
}

func digitsOfNegativeInt32() []int {
	_, min := maxMinInt32()
	result := make([]int, 0)
	for min < 0 {
		result = append(result, int((min%10)*-1))
		min = min / 10
	}
	return result
}

func maxMinInt32() (int32, int32) {
	const MaxUint32 = ^uint32(0)
	return int32(MaxUint32 >> 1), -int32(MaxUint32>>1) - 1
}
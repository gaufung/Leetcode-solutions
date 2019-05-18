/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
func longestPalindrome(s string) string {
	length, result := 0, make([]rune, 0)
	buffer := []rune(s)
	for idx, _ := range s {
		current_length := palindromeLengthSymmetic(s, idx, idx+1)
		if current_length > length {
			result = buffer[idx-(current_length/2)+1 : idx+(current_length/2)+1]
			length = current_length
		}
		current_length = palindromeLengthNonSymmetric(s, idx)
		if current_length > length {
			result = buffer[idx-(current_length-1)/2 : idx+(current_length-1)/2+1]
			length = current_length
		}
	}
	return string(result)
}
func palindromeLengthSymmetic(s string, left, right int) int {
	result := 0
	for left > -1 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
			result += 2
		} else {
			break
		}
	}
	return result
}

func palindromeLengthNonSymmetric(s string, middle int) int {
	result := 1
	return result + palindromeLengthSymmetic(s, middle-1, middle+1)
}


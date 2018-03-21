import "strings"
var INT_MAX = 2147483647
var INT_MIN = -2147483648

func myAtoi(str string) int {
	if "" == strings.Trim(str, " "){
		return 0
	}
	if digits, flag := digits(str); flag{
		//positive
		result := 0
		for _, digit := range(digits){
			result = result * 10 + digit
			if result > INT_MAX{
				return INT_MAX
			}
		}
		return result

	}else{
		//negative
		result := 0
		for _, digit := range digits {
			result = result * 10 - digit
			if result < INT_MIN {
				return INT_MIN
			}
		}
		return result

	}

}

func digits(str string) ([]int, bool) {
	str = strings.Trim(str, " ")
	result := make([]int, 0)
	flag := true
	if str[0] == '-' {
		flag = false
		str = str[1:]
	}else if str[0] == '+'{
		flag = true
		str = str[1:]
	}
	for _, ch := range str {
		if digit, ok := digit(ch); ok {
			result = append(result, digit)
		}else{
			break
		}
	}
	return result, flag
}

func digit(ch rune) (int, bool) {
	if ch >= '0' && ch <= '9' {
		return int(ch - '0'), true
	}else{
		return 0, false
	}
}
import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 */
func readBinaryWatch(num int) []string {
	result := []string{}
	for i := 0; i <= 4; i++ {
		j := num - i
		if j >= 0 {
			hours := hour(i)
			minutes := minute(j)
			for _, h := range hours {
				for _, m := range minutes {
					result = append(result, h+":"+m)
				}
			}
		}
	}
	return result
}

func hour(nums int) []string {
	candidates := counter(nums, 4)
	result := make([]string, 0)
	for _, val := range candidates {
		if val < 12 {
			result = append(result, strconv.Itoa(val))
		}
	}
	return result
}

func minute(nums int) []string {
	candidates := counter(nums, 6)
	result := make([]string, 0)
	for _, val := range candidates {
		if val < 60 {
			result = append(result, fmt.Sprintf("%02d", val))
		}
	}
	return result
}

func counter(nums, count int) []int {
	if nums == 0 {
		return []int{0}
	}
	if nums >= count {
		return []int{1<<uint(count) - 1}
	}
	result := make([]int, 0)
	leftCount := counter(nums, count-1)
	for _, val := range leftCount {
		result = append(result, val)
	}
	leftCount = counter(nums-1, count-1)
	for _, val := range leftCount {
		result = append(result, val+(1<<uint(count-1)))
	}
	return result
}



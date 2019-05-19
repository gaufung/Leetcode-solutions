import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 *
 * https://leetcode-cn.com/problems/binary-watch/description/
 *
 * algorithms
 * Easy (47.12%)
 * Likes:    60
 * Dislikes: 0
 * Total Accepted:    3.6K
 * Total Submissions: 7.7K
 * Testcase Example:  '0'
 *
 * 二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。
 *
 * 每个 LED 代表一个 0 或 1，最低位在右侧。
 *
 *
 *
 * 例如，上面的二进制手表读取 “3:25”。
 *
 * 给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。
 *
 * 案例:
 *
 *
 * 输入: n = 1
 * 返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16",
 * "0:32"]
 *
 *
 *
 * 注意事项:
 *
 *
 * 输出的顺序没有要求。
 * 小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
 * 分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。
 *
 *
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


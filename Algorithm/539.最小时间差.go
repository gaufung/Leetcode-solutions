import (
	"sort"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=539 lang=golang
 *
 * [539] 最小时间差
 *
 * https://leetcode-cn.com/problems/minimum-time-difference/description/
 *
 * algorithms
 * Medium (48.30%)
 * Likes:    26
 * Dislikes: 0
 * Total Accepted:    1.7K
 * Total Submissions: 3.4K
 * Testcase Example:  '["23:59","00:00"]'
 *
 * 给定一个 24 小时制（小时:分钟）的时间列表，找出列表中任意两个时间的最小时间差并已分钟数表示。
 *
 *
 * 示例 1：
 *
 *
 * 输入: ["23:59","00:00"]
 * 输出: 1
 *
 *
 *
 * 备注:
 *
 *
 * 列表中时间数在 2~20000 之间。
 * 每个时间取值在 00:00~23:59 之间。
 *
 *
 */
func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	result := timeElapse(timePoints[1], timePoints[0])
	for i := 1; i < len(timePoints); i++ {
		val := timeElapse(timePoints[i], timePoints[i-1])
		result = min(result, val)
	}
	val := timeElapse(timePoints[0], timePoints[len(timePoints)-1])
	result = min(result, val)
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func timeElapse(t1, t2 string) int {
	if t1 == t2 {
		return 0
	}
	result := 0
	if strings.Compare(t1, t2) == 1 {
		t1Hour, t1Minute := format(t1)
		t2Hour, t2Minute := format(t2)

		if t1Minute >= t2Minute {
			result = t1Minute - t2Minute + 60*(t1Hour-t2Hour)
		} else {
			result = 60 + t1Minute - t2Minute + 60*(t1Hour-1-t2Hour)
		}
	} else {
		t1Hour, t1Minute := format(t1)
		t2Hour, t2Minute := format(t2)
		t1Hour += 24
		if t1Minute >= t2Minute {
			result = t1Minute - t2Minute + 60*(t1Hour-t2Hour)
		} else {
			result = 60 + t1Minute - t2Minute + 60*(t1Hour-1-t2Hour)
		}
	}
	return result
}

func format(t string) (int, int) {
	index := strings.Index(t, ":")
	hour, _ := strconv.Atoi(t[:index])
	minute, _ := strconv.Atoi(t[index+1:])
	return hour, minute
}


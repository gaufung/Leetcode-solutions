/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 *
 * https://leetcode-cn.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (45.31%)
 * Likes:    85
 * Dislikes: 0
 * Total Accepted:    4.6K
 * Total Submissions: 10K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 现在你总共有 n 门课需要选，记为 0 到 n-1。
 *
 * 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
 *
 * 给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？
 *
 * 示例 1:
 *
 * 输入: 2, [[1,0]]
 * 输出: true
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
 *
 * 示例 2:
 *
 * 输入: 2, [[1,0],[0,1]]
 * 输出: false
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
 *
 * 说明:
 *
 *
 * 输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
 * 你可以假定输入的先决条件中没有重复的边。
 *
 *
 * 提示:
 *
 *
 * 这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
 * 通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
 *
 * 拓扑排序也可以通过 BFS 完成。
 *
 *
 *
 */
func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 || len(prerequisites) <= 0 {
		return true
	}
	m := make(map[int][]int, 0)
	notStart := make([]bool, numCourses)
	for _, neighbour := range prerequisites {
		from, to := neighbour[1], neighbour[0]
		if val, ok := m[from]; ok {
			m[from] = append(val, to)
		} else {
			m[from] = []int{to}
		}
		notStart[to] = true
	}
	in, out := make([]int, numCourses), make([]int, numCourses)
	visited := make([]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		if notStart[i] == false {
			index = 0
			if DFS(i, m, visited, in, out) == false {
				return false
			}
		}
	}
	for _, val := range visited {
		if val == false {
			return false
		}
	}
	return true
}

var index int

func DFS(seed int, m map[int][]int, visited []bool, in, out []int) bool {
	neighbours := m[seed]
	if neighbours == nil {
		neighbours = []int{}
	}
	index++
	in[seed] = index
	visited[seed] = true
	for _, v := range neighbours {
		if visited[v] == false {
			if DFS(v, m, visited, in, out) == false {
				return false
			}
		} else {
			if in[v] == 0 || out[v] == 0 {
				return false
			}
		}
	}
	index++
	out[seed] = index
	return true
}


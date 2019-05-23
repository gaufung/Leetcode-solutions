/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] 最小高度树
 *
 * https://leetcode-cn.com/problems/minimum-height-trees/description/
 *
 * algorithms
 * Medium (34.08%)
 * Likes:    37
 * Dislikes: 0
 * Total Accepted:    1.3K
 * Total Submissions: 3.9K
 * Testcase Example:  '4\n[[1,0],[1,2],[1,3]]'
 *
 *
 * 对于一个具有树特征的无向图，我们可选择任何一个节点作为根。图因此可以成为树，在所有可能的树中，具有最小高度的树被称为最小高度树。给出这样的一个图，写出一个函数找到所有的最小高度树并返回他们的根节点。
 *
 * 格式
 *
 * 该图包含 n 个节点，标记为 0 到 n - 1。给定数字 n 和一个无向边 edges 列表（每一个边都是一对标签）。
 *
 * 你可以假设没有重复的边会出现在 edges 中。由于所有的边都是无向边， [0, 1]和 [1, 0] 是相同的，因此不会同时出现在 edges 里。
 *
 * 示例 1:
 *
 * 输入: n = 4, edges = [[1, 0], [1, 2], [1, 3]]
 *
 * ⁠       0
 * ⁠       |
 * ⁠       1
 * ⁠      / \
 * ⁠     2   3
 *
 * 输出: [1]
 *
 *
 * 示例 2:
 *
 * 输入: n = 6, edges = [[0, 3], [1, 3], [2, 3], [4, 3], [5, 4]]
 *
 * ⁠    0  1  2
 * ⁠     \ | /
 * ⁠       3
 * ⁠       |
 * ⁠       4
 * ⁠       |
 * ⁠       5
 *
 * 输出: [3, 4]
 *
 * 说明:
 *
 *
 * 根据树的定义，树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。
 * 树的高度是指根节点和叶子节点之间最长向下路径上边的数量。
 *
 *
 */
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	heights := make([]int, n)
	leaves := findLeaves(edges)
	lowheight := n
	// for leaf, _ := range leaves {
	// 	height := bfs(leaf, edges, n, leaves)

	// }
	isfound := false
	for i := 0; i < n; i++ {
		if _, ok := leaves[i]; !ok {
			isfound = true
			height := bfs(i, edges, n, leaves)
			if height < lowheight {
				lowheight = height
			}
			heights[i] = height
		}

	}
	if !isfound {
		lowheight = 0
	}
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if heights[i] == lowheight {
			result = append(result, i)
		}
	}
	return result
}

func findLeaves(edges [][]int) map[int]int {
	m := make(map[int]int, 0)
	for _, edge := range edges {
		if _, ok := m[edge[0]]; ok {
			m[edge[0]] += 1
		} else {
			m[edge[0]] = 1
		}
		if _, ok := m[edge[1]]; ok {
			m[edge[1]] += 1
		} else {
			m[edge[1]] = 1
		}
	}
	for k, v := range m {
		if v >= 2 {
			delete(m, k)
		}
	}
	return m
}

func bfs(seed int, edges [][]int, n int, leaves map[int]int) int {
	visited := make([]bool, n)
	heights := make([]int, n)
	heights[seed] = 0
	visited[seed] = true
	queue := []int{seed}
	path := 1
	for len(queue) > 0 {
		newQueue := make([]int, 0)
		for _, q := range queue {
			for _, edge := range edges {
				if edge[0] == q && visited[edge[1]] == false {
					visited[edge[1]] = true
					newQueue = append(newQueue, edge[1])
					heights[edge[1]] = path
				}
				if edge[1] == q && visited[edge[0]] == false {
					visited[edge[0]] = true
					newQueue = append(newQueue, edge[0])
					heights[edge[0]] = path
				}
			}
		}
		path++
		queue = newQueue
	}
	maxHeight := 0
	for i := 0; i < n; i++ {
		if _, ok := leaves[i]; ok {
			if heights[i] > maxHeight {
				maxHeight = heights[i]
			}
		}
	}
	return maxHeight
}




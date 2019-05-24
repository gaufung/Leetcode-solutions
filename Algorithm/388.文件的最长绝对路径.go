/*
 * @lc app=leetcode.cn id=388 lang=golang
 *
 * [388] 文件的最长绝对路径
 *
 * https://leetcode-cn.com/problems/longest-absolute-file-path/description/
 *
 * algorithms
 * Medium (43.87%)
 * Likes:    7
 * Dislikes: 0
 * Total Accepted:    422
 * Total Submissions: 962
 * Testcase Example:  '"dir\\n\\tsubdir1\\n\\tsubdir2\\n\\t\\tfile.ext"'
 *
 * 假设我们以下述方式将我们的文件系统抽象成一个字符串:
 *
 * 字符串 "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext" 表示:
 *
 *
 * dir
 * ⁠   subdir1
 * ⁠   subdir2
 * ⁠       file.ext
 *
 *
 * 目录 dir 包含一个空的子目录 subdir1 和一个包含一个文件 file.ext 的子目录 subdir2 。
 *
 * 字符串
 * "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
 * 表示:
 *
 *
 * dir
 * ⁠   subdir1
 * ⁠       file1.ext
 * ⁠       subsubdir1
 * ⁠   subdir2
 * ⁠       subsubdir2
 * ⁠           file2.ext
 *
 *
 * 目录 dir 包含两个子目录 subdir1 和 subdir2。 subdir1 包含一个文件 file1.ext 和一个空的二级子目录
 * subsubdir1。subdir2 包含一个二级子目录 subsubdir2 ，其中包含一个文件 file2.ext。
 *
 * 我们致力于寻找我们文件系统中文件的最长 (按字符的数量统计) 绝对路径。例如，在上述的第二个例子中，最长路径为
 * "dir/subdir2/subsubdir2/file2.ext"，其长度为 32 (不包含双引号)。
 *
 * 给定一个以上述格式表示文件系统的字符串，返回文件系统中文件的最长绝对路径的长度。 如果系统中没有文件，返回 0。
 *
 * 说明:
 *
 *
 * 文件名至少存在一个 . 和一个扩展名。
 * 目录或者子目录的名字不能包含 .。
 *
 *
 * 要求时间复杂度为 O(n) ，其中 n 是输入字符串的大小。
 *
 * 请注意，如果存在路径 aaaaaaaaaaaaaaaaaaaaa/sth.png 的话，那么  a/aa/aaa/file1.txt
 * 就不是一个最长的路径。
 *
 */
func lengthLongestPath(input string) int {
	stack := NewStack()
	topIndent := 0
	index := 0
	result := 0
	input = input + "\n"
	for index < len(input) {
		j := index
		indent := 0
		isFile := false
		for input[j] != '\n' {
			if input[j] == '\t' {
				indent++
			}
			if input[j] == '.' {
				isFile = true
			}
			j++
		}
		if stack.Size() > 0 {
			for k := topIndent; k >= indent; k-- {
				stack.Pop()
			}
		}
		name := input[index+indent : j]
		stack.Push(name)
		topIndent = indent
		if isFile {
			result = max(result, pathSize(stack))
		}
		index = j + 1
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pathSize(s *Stack) int {
	if s.Size() == 0 {
		return 0
	}
	val := s.Size() - 1
	for i := 0; i < s.Size(); i++ {
		val += len(s.elements[i])
	}
	return val
}

type Stack struct {
	elements []string
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]string, 0),
	}
}
func (s *Stack) Size() int {
	return len(s.elements)
}
func (s *Stack) Push(val string) {
	s.elements = append(s.elements, val)
}
func (s *Stack) Pop() string {
	val := s.elements[s.Size()-1]
	s.elements = s.elements[:s.Size()-1]
	return val
}
func (s *Stack) Top() string {
	return s.elements[s.Size()-1]
}

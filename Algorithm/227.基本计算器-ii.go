/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 *
 * https://leetcode-cn.com/problems/basic-calculator-ii/description/
 *
 * algorithms
 * Medium (31.39%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    3K
 * Total Submissions: 9.7K
 * Testcase Example:  '"3+2*2"'
 *
 * 实现一个基本的计算器来计算一个简单的字符串表达式的值。
 *
 * 字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
 *
 * 示例 1:
 *
 * 输入: "3+2*2"
 * 输出: 7
 *
 *
 * 示例 2:
 *
 * 输入: " 3/2 "
 * 输出: 1
 *
 * 示例 3:
 *
 * 输入: " 3+5 / 2 "
 * 输出: 5
 *
 *
 * 说明：
 *
 *
 * 你可以假设所给定的表达式都是有效的。
 * 请不要使用内置的库函数 eval。
 *
 *
 */
func calculate(s string) int {
	numbersStack := NewStack()
	operatorStack := NewStack()
	lexer := NewLexer(s)
	operations := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	precedences := map[string]map[string]bool{
		"+": map[string]bool{"+": false, "-": false, "*": false, "/": false},
		"-": map[string]bool{"+": false, "-": false, "*": false, "/": false},
		"*": map[string]bool{"+": true, "-": true, "*": false, "/": false},
		"/": map[string]bool{"+": true, "-": true, "*": false, "/": false},
	}
	for {
		token, literal := lexer.Next()
		if token == EOF {
			break
		}
		if token == NUMBER {
			number := atoi(literal)
			numbersStack.Push(number)
		} else {
			for {
				if operatorStack.Empty() {
					operatorStack.Push(literal)
					break
				} else {
					top := operatorStack.Top().(string)
					if precedences[literal][top] == false {
						argument2 := numbersStack.Pop().(int)
						argument1 := numbersStack.Pop().(int)
						result := operations[top](argument1, argument2)
						operatorStack.Pop()
						numbersStack.Push(result)
					} else {
						operatorStack.Push(literal)
						break
					}
				}
			}

		}
	}
	for !operatorStack.Empty() {
		operator := operatorStack.Pop().(string)

		argument2 := numbersStack.Pop().(int)
		argument1 := numbersStack.Pop().(int)
		result := operations[operator](argument1, argument2)
		numbersStack.Push(result)
	}
	return numbersStack.Top().(int)
}

type TYPE int

var NUMBER TYPE = 0
var OPERATOR TYPE = 1
var EOF TYPE = 3

type Lexer struct {
	input   []byte
	pos     int
	readPos int
}

func NewLexer(expression string) *Lexer {
	return &Lexer{
		input:   []byte(expression),
		pos:     0,
		readPos: 0,
	}
}

func (l *Lexer) Next() (TYPE, string) {
	l.skipWhiteSpace()
	if l.pos == len(l.input) {
		return EOF, ""
	}
	if l.input[l.pos] == '+' || l.input[l.pos] == '-' || l.input[l.pos] == '*' || l.input[l.pos] == '/' {
		literal := string(l.input[l.pos : l.pos+1])
		l.pos++
		return OPERATOR, literal
	} else {
		return l.readNumber()
	}
}

func (l *Lexer) skipWhiteSpace() {
	for l.pos < len(l.input) && l.input[l.pos] == ' ' {
		l.pos++
	}
}

func (l *Lexer) readNumber() (TYPE, string) {
	l.readPos = l.pos + 1
	for l.readPos < len(l.input) && isDigits(l.input[l.readPos]) {
		l.readPos++
	}
	literal := string(l.input[l.pos:l.readPos])
	l.pos = l.readPos
	return NUMBER, literal
}

func isDigits(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func atoi(str string) int {
	val := 0
	for i := 0; i < len(str); i++ {
		val = val*10 + int(str[i]-'0')
	}
	return val
}

type Stack struct {
	elements []interface{}
	size     int
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
		size:     0,
	}
}

func (s *Stack) Push(v interface{}) {
	s.elements = append(s.elements, v)
	s.size++
}

func (s *Stack) Pop() interface{} {
	s.size--
	val := s.elements[s.size]
	s.elements = s.elements[:s.size]
	return val
}
func (s *Stack) Top() interface{} {
	return s.elements[s.size-1]
}

func (s *Stack) Empty() bool {
	return s.size == 0
}


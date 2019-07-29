import "strconv"

/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 *
 * https://leetcode-cn.com/problems/basic-calculator/description/
 *
 * algorithms
 * Hard (32.76%)
 * Likes:    78
 * Dislikes: 0
 * Total Accepted:    2.9K
 * Total Submissions: 8.8K
 * Testcase Example:  '"1 + 1"'
 *
 * 实现一个基本的计算器来计算一个简单的字符串表达式的值。
 *
 * 字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。
 *
 * 示例 1:
 *
 * 输入: "1 + 1"
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: " 2-1 + 2 "
 * 输出: 3
 *
 * 示例 3:
 *
 * 输入: "(1+(4+5+2)-3)+(6+8)"
 * 输出: 23
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
	lexer := NewLexer(s)
	operator := NewStack()
	operand := NewStack()
	for {
		token := lexer.NextToken()
		if token.Type == EOF {
			break
		}
		if token.Type == LeftPare {
			operator.Push(token)
		}
		if token.Type == RightPare {
			for operator.Top().Type != LeftPare {
				op := operator.Pop()
				b := operand.Pop()
				a := operand.Pop()
				operand.Push(calc(a, b, op))
			}
			operator.Pop()
			for operator.Size() > 0 && operator.Top().Type != LeftPare && operand.Size() > 1 {
				op := operator.Pop()
				b := operand.Pop()
				a := operand.Pop()
				operand.Push(calc(a, b, op))
			}
		}
		if token.Type == Operator {
			operator.Push(token)
		}
		if token.Type == Number {
			if operator.Size() == 0 || operator.Top().Type == LeftPare || operand.Size() == 0 {
				operand.Push(token)
			} else {
				a := operand.Pop()
				op := operator.Pop()
				operand.Push(calc(a, token, op))
				for operator.Size() > 0 && operator.Top().Type != LeftPare && operand.Size() > 1 {
					op := operator.Pop()
					b := operand.Pop()
					a := operand.Pop()
					operand.Push(calc(a, b, op))
				}
			}
		}

	}
	for operator.Size() != 0 {
		b := operand.Pop()
		a := operand.Pop()
		op := operator.Pop()
		operand.Push(calc(a, b, op))
	}
	val, _ := strconv.Atoi(operand.Top().Literal)
	return val
}

func calc(t1, t2 Token, op Token) Token {
	a, _ := strconv.Atoi(t1.Literal)
	b, _ := strconv.Atoi(t2.Literal)
	if op.Literal == "-" {
		return Token{Type: Number, Literal: strconv.Itoa(a - b)}
	} else {
		return Token{Type: Number, Literal: strconv.Itoa(a + b)}
	}
}

type TokenType string

const (
	Number    TokenType = "Number"
	LeftPare  TokenType = "LeftPare"
	RightPare TokenType = "RightPare"
	Operator  TokenType = "Operator"
	EOF       TokenType = "EOF"
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	position     int
	nextPosition int
	input        string
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		position:     0,
		nextPosition: 1,
		input:        input,
	}
}

func (l *Lexer) NextToken() Token {
	for l.position < len(l.input) {
		if l.input[l.position] == ' ' {
			l.skipWhite()
			continue
		}
		switch l.input[l.position] {
		case '(':
			l.position++
			l.nextPosition++
			return Token{Type: LeftPare, Literal: "("}
		case ')':
			l.position++
			l.nextPosition++
			return Token{Type: RightPare, Literal: ")"}
		case '+':
			l.position++
			l.nextPosition++
			return Token{Type: Operator, Literal: "+"}
		case '-':
			l.position++
			l.nextPosition++
			return Token{Type: Operator, Literal: "-"}
		default:
			return l.readNumber()
		}
	}
	return Token{Type: EOF, Literal: ""}
}

func (l *Lexer) skipWhite() {
	for l.position < len(l.input) && l.input[l.position] == ' ' {
		l.position++
		l.nextPosition++
	}
}

func (l *Lexer) readNumber() Token {
	for l.nextPosition < len(l.input) && isDigit(l.input[l.nextPosition]) {
		l.nextPosition++
	}
	number := l.input[l.position:l.nextPosition]
	l.position = l.nextPosition
	l.nextPosition++
	return Token{Type: Number, Literal: number}
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

type Stack struct {
	elements []Token
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]Token, 0),
	}
}

func (s *Stack) Push(val Token) {
	s.elements = append(s.elements, val)
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Pop() Token {

	val := s.elements[s.Size()-1]
	s.elements = s.elements[:s.Size()-1]
	return val
}

func (s *Stack) Top() Token {
	return s.elements[s.Size()-1]
}
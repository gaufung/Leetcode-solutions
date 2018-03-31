func isValid(s string) bool {
    ss := newStack()
    for _, ch := range s {
        if ss.isEmpty() {
            ss.push(ch)
        }else{
            if isMatch(ss.peek(), ch){
                ss.pop()
            }else{
                ss.push(ch)
            }
        }
    }
    return ss.isEmpty()
}

func isMatch(left, right rune) bool {
    return (left=='(' && right==')') || (left=='{' && right=='}') || (left=='[' && right==']')
}

type stack struct{
    values []rune
    pos int
}

func newStack() *stack{
    return &stack{make([]rune, 1), -1}
}

func (s *stack) pop() rune{
    val := s.values[s.pos]
    s.pos --
    return val
}

func (s *stack) push(val rune){
    if s.pos == len(s.values) - 1 {
        backup := make([]rune, 2 * len(s.values))
        copy(backup, s.values)
        s.values = backup
    }
    s.pos ++
    s.values[s.pos] = val
}

func (s *stack) isEmpty() bool {
    return s.pos == -1
}
func (s *stack) peek() rune {
    return s.values[s.pos]
}
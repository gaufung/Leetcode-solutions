func longestCommonPrefix(strs []string) string {
	tree := NewTriTree()
	for _, str := range strs {
		if str == ""{
			return ""
		}else{
			tree.InsertString(str)
		}
	}
	return tree.LongestPrefix(len(strs))
}

type Node struct {
	Value byte
	Frequence int
	Nexts map[byte]*Node
}

func EmptyNode() *Node{
	return &Node{0,0,make(map[byte]*Node)}
}

func ValueNode(val byte) *Node{
	return &Node{val, 1, make(map[byte]*Node)}
}

type TriTree struct {
	Root *Node
}

func NewTriTree() *TriTree {
	return &TriTree{EmptyNode()}
}

func (tree *TriTree) InsertString(str string) {
	pointer := tree.Root
	for i:=0; i<len(str); i++ {
		if node, ok := pointer.Nexts[str[i]]; ok {
			pointer = node
			pointer.Frequence ++
		}else{
			pointer.Nexts[str[i]] = ValueNode(str[i])
			pointer = pointer.Nexts[str[i]]
		}
	}
}




func (tree *TriTree) LongestPrefix(expect int) string {
	result := make([]byte,0)
	pointer := tree.Root
	flag := true
	for flag {
		dictionary := pointer.Nexts
		flag = false
		for char, node := range dictionary {
			fmt.Println(node.Value, node.Frequence)
			if node.Frequence == expect{
				result = append(result, char)
				pointer = node
				flag = true
			}
		}
	}
	return string(result)
}
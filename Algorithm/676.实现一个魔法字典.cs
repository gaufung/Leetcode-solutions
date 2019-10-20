/*
 * @lc app=leetcode.cn id=676 lang=csharp
 *
 * [676] 实现一个魔法字典
 *
 * https://leetcode-cn.com/problems/implement-magic-dictionary/description/
 *
 * algorithms
 * Medium (53.78%)
 * Likes:    16
 * Dislikes: 0
 * Total Accepted:    1K
 * Total Submissions: 1.9K
 * Testcase Example:  '["MagicDictionary", "buildDict", "search", "search", "search", "search"]\n[[], [["hello","leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]'
 *
 * 实现一个带有buildDict, 以及 search方法的魔法字典。
 * 
 * 对于buildDict方法，你将被给定一串不重复的单词来构建一个字典。
 * 
 * 对于search方法，你将被给定一个单词，并且判定能否只将这个单词中一个字母换成另一个字母，使得所形成的新单词存在于你构建的字典中。
 * 
 * 示例 1:
 * 
 * 
 * Input: buildDict(["hello", "leetcode"]), Output: Null
 * Input: search("hello"), Output: False
 * Input: search("hhllo"), Output: True
 * Input: search("hell"), Output: False
 * Input: search("leetcoded"), Output: False
 * 
 * 
 * 注意:
 * 
 * 
 * 你可以假设所有输入都是小写字母 a-z。
 * 为了便于竞赛，测试所用的数据量很小。你可以在竞赛结束后，考虑更高效的算法。
 * 请记住重置MagicDictionary类中声明的类变量，因为静态/类变量会在多个测试用例中保留。 请参阅这里了解更多详情。
 * 
 * 
 */

// @lc code=start
public class MagicDictionary {

    private Node root;
    /** Initialize your data structure here. */
    public MagicDictionary() {
        this.root = new Node(' ');   
    }
    
    /** Build a dictionary through a list of words */
    public void BuildDict(string[] dict) {
        foreach (string word in dict)
        {
            this.Insert(word);
        }
    }

    private void Insert(string word)
    {
        Node cur = this.root;
        foreach (char ch in word)
        {
            int idx = (int)(ch - 'a');
            if(cur.Next[idx] == null)
            {
                cur.Next[idx] = new Node(ch);
            }
            cur = cur.Next[idx];
        }
        cur.IsTerminate = true;
    }
    
    /** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
    public bool Search(string word) {
        Node cur = this.root; 
        bool found = false;
        for(int i = 0; i < word.Length; i++)
        {
            int idx = (int)(word[i] - 'a');
            for(int j = 0; j < 26; j++)
            {
                if(cur.Next[j] != null && j != idx)
                {
                    found |= this.Search(cur.Next[j], word.Substring(i+1));
                }
            }
            if(found)
            {
                return true;
            }
            else
            {
                if(cur.Next[idx]!=null)
                {
                    cur = cur.Next[idx];
                }
                else
                {
                    return false;
                }
            }
        }
        return false;
    }

    private bool Search(Node node, string subWord)
    {
        foreach (var ch in subWord)
        {
            if(node == null)
            {
                return false;
            }
            int idx = (int)(ch - 'a');
            if(node.Next[idx] == null)
            {
                return false;
            }
            else
            {
                node = node.Next[idx];
            }
        }
        return node.IsTerminate;

    }
}

public class Node 
{
    public char Ch;
    public Node[] Next;

    public bool IsTerminate;

    public Node(char ch)
    {
        this.Ch = ch;
        this.Next = new Node[26];
    }
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * MagicDictionary obj = new MagicDictionary();
 * obj.BuildDict(dict);
 * bool param_2 = obj.Search(word);
 */
// @lc code=end


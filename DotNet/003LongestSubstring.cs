public class Solution {
    public int LengthOfLongestSubstring(string s) {
        int result = 0;
        IDictionary<char, int> map = new Dictionary<char, int>();
        for(int i=0; i< s.Length;i++){
            if(map.ContainsKey(s[i])){
                result = Math.Max(result, map.Count);
                int start = map[s[i]]+1;
                map = new Dictionary<char, int>();
                for(int j = start; j<=i;j++){
                    map.Add(s[j], j);
                }
            }else{
                map.Add(s[i], i);
            }
        }
        result = Math.Max(result, map.Count);
        return result;
    }
}
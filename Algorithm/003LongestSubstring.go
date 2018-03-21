func max(x, y int) int {
    if x>y {
        return x
    }else{
        return y
    }
}
func lengthOfLongestSubstring(s string) int {
    result := 0
    maps := make(map[rune] int)
    for idx, char:= range(s) {
        if start, ok := maps[char]; ok {
            result = max(result, len(maps))
            maps = make(map[rune] int)
            for ii, val := range(s[start:idx+1]){
                maps[val] = start + ii
            }
        }else{
            maps[char] = idx
        }
    }
    result = max(result, len(maps))
    return result
    
}
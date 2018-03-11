class Solution:
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        result = 0
        maps = {}
        for idx in range(len(s)):
            char = s[idx]
            if char in maps:
                result = max(result, len(maps))
                start = maps[char]
                maps = {}
                for ii in range(start, idx+1):
                    maps[s[ii]]=ii
            else:
                maps[char]=idx
        result = max(result, len(maps))
        return result
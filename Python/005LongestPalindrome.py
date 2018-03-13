class Solution:
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        length  = 0
        result = ''
        for i in range(len(s)):
            current_length = self.palindromeLengthSymmetric(s, i,i+1)
            if current_length > length:
                result = s[i-int(current_length/2)+1:i+int(current_length/2)+1]
                length = current_length
            current_length = self.palindromeLengthNonSymmeteric(s,i)
            if current_length > length:
                result = s[i-int((current_length-1)/2):i+int((current_length-1)/2+1)]
                length = current_length
        return result
        
    def palindromeLengthSymmetric(self, s, left, right):
        result = 0
        while left > -1 and right < len(s):
            if s[left] == s[right] :
                left -= 1
                right += 1
                result += 2
            else:
                break
        return result
    
    def palindromeLengthNonSymmeteric(self, s, middle):
        result = 1
        return result + self.palindromeLengthSymmetric(s, middle-1, middle+1)
public class Solution {
    public string LongestPalindrome(string s) {
        int length = 0;
        string palindrome="";
        for(int i = 0; i<s.Length; i++){
            int currentLength = PalindromicLength(s, i, i+1);
            if (currentLength > length){
                palindrome=s.Substring(i-currentLength/2+1,currentLength);
                length = currentLength;
            }
            currentLength = PalindromicLength(s, i);
            if(currentLength > length){
                palindrome = s.Substring(i-(currentLength-1)/2, currentLength);
                length = currentLength;
            }
        }
        return palindrome;
        
    }
    private int PalindromicLength(string s, int left, int right){
        int result = 0;
        while(left > -1 && right < s.Length){
            if(s[left]==s[right]){
                left--;
                right++;
                result += 2;
            }else{
                break;
            }
        }
        return result;
    }
    private int PalindromicLength(string s, int middle){
        int result = 1;
        return result + PalindromicLength(s, middle-1, middle+1);
    }
}
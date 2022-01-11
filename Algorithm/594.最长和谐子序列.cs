/*
 * @lc app=leetcode.cn id=594 lang=csharp
 *
 * [594] 最长和谐子序列
 *
 * https://leetcode-cn.com/problems/longest-harmonious-subsequence/description/
 *
 * algorithms
 * Easy (56.60%)
 * Likes:    302
 * Dislikes: 0
 * Total Accepted:    59.9K
 * Total Submissions: 105.9K
 * Testcase Example:  '[1,3,2,2,5,2,3,7]'
 *
 * 和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。
 * 
 * 现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。
 * 
 * 数组的子序列是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,3,2,2,5,2,3,7]
 * 输出：5
 * 解释：最长的和谐子序列是 [3,2,2,2,3]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,3,4]
 * 输出：2
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1,1,1,1]
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -10^9 
 * 
 * 
 */

// @lc code=start
public class Solution {
    public int FindLHS(int[] nums) {
        Array.Sort(nums);
        int cnt = 0;
        for(int i = 0; i < nums.Length; i++) {
            int val = nums[i];
            int j = i;
            while (j < nums.Length) {
                if (nums[j] == nums[i]) {
                    j++;
                    continue;
                }
                if (nums[j] == nums[i] + 1) {
                    j++;
                    continue;
                }
                if (nums[j] > nums[i] + 1) {
                    break;
                }
            }
            if (nums[j-1] == nums[i] + 1){
                var length = j - i;
            if (length > cnt) {
                cnt = length;
            }
            }

            
        }

        return cnt;
    }
}
// @lc code=end


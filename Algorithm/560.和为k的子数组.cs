/*
 * @lc app=leetcode.cn id=560 lang=csharp
 *
 * [560] 和为K的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (41.55%)
 * Likes:    151
 * Dislikes: 0
 * Total Accepted:    9.2K
 * Total Submissions: 22K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
 * 
 * 示例 1 :
 * 
 * 
 * 输入:nums = [1,1,1], k = 2
 * 输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
 * 
 * 
 * 说明 :
 * 
 * 
 * 数组的长度为 [1, 20,000]。
 * 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
 * 
 * 
 */

// @lc code=start
public class Solution {
    public int SubarraySum(int[] nums, int k) {
        int N = nums.Length;
        //int[,] mat = new int[N,N];
        int count = 0;
        int sum = 0;
        for(int i = 0; i < N; i++)
        {
            for (int j = i; j < N; j++)
            {
                if(i == j)
                {
                    sum = nums[i];
                }
                else
                {
                    sum += nums[j];
                }
                if(sum == k)
                {
                    count++;
                }
            }
        }
        return count;


    }
}
// @lc code=end


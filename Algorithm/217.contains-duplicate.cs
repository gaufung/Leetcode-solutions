/*
 * @lc app=leetcode id=217 lang=csharp
 *
 * [217] Contains Duplicate
 *
 * https://leetcode.com/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (59.61%)
 * Likes:    3252
 * Dislikes: 905
 * Total Accepted:    1.2M
 * Total Submissions: 2M
 * Testcase Example:  '[1,2,3,1]'
 *
 * Given an integer array nums, return true if any value appears at least twice
 * in the array, and return false if every element is distinct.
 * 
 * 
 * Example 1:
 * Input: nums = [1,2,3,1]
 * Output: true
 * Example 2:
 * Input: nums = [1,2,3,4]
 * Output: false
 * Example 3:
 * Input: nums = [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * 
 * 
 */

// @lc code=start
public class Solution {
    public bool ContainsDuplicate(int[] nums) {
        Array.Sort(nums);
        for(int i = 0; i < nums.Length - 1; i++)
        {
            if (nums[i] == nums[i+1])
            {
                return true;
            }
        }
        return false;
    }
}
// @lc code=end


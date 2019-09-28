/*
 * @lc app=leetcode.cn id=836 lang=csharp
 *
 * [836] 矩形重叠
 *
 * https://leetcode-cn.com/problems/rectangle-overlap/description/
 *
 * algorithms
 * Easy (41.47%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    3.5K
 * Total Submissions: 8.3K
 * Testcase Example:  '[0,0,2,2]\n[1,1,3,3]'
 *
 * 矩形以列表 [x1, y1, x2, y2] 的形式表示，其中 (x1, y1) 为左下角的坐标，(x2, y2) 是右上角的坐标。
 * 
 * 如果相交的面积为正，则称两矩形重叠。需要明确的是，只在角或边接触的两个矩形不构成重叠。
 * 
 * 给出两个矩形，判断它们是否重叠并返回结果。
 * 
 * 示例 1：
 * 
 * 输入：rec1 = [0,0,2,2], rec2 = [1,1,3,3]
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 输入：rec1 = [0,0,1,1], rec2 = [1,0,2,1]
 * 输出：false
 * 
 * 
 * 说明：
 * 
 * 
 * 两个矩形 rec1 和 rec2 都以含有四个整数的列表的形式给出。
 * 矩形中的所有坐标都处于 -10^9 和 10^9 之间。
 * 
 * 
 */
public class Solution {
    public bool IsRectangleOverlap(int[] rec1, int[] rec2) {
        int[] leftBottom = new int[] {rec1[0], rec1[1]};
        int[] leftTop = new int[] {rec1[0], rec1[3]};
        int[] rightBottom = new int[] { rec1[2], rec1[1]};
        int[] rightTop = new int[] { rec1[2], rec1[3]};
        var result = IsPointInRectangel(leftBottom, rec2) || IsPointInRectangel(leftTop, rec2) || IsPointInRectangel(rightBottom, rec2) || IsPointInRectangel(rightTop, rec2);
        if(result)
        {
            return true;
        }
        leftBottom = new int[] {rec2[0], rec2[1]};
        leftTop = new int[] {rec2[0], rec2[3]};
        rightBottom = new int[] { rec2[2], rec2[1]};
        rightTop = new int[] { rec2[2], rec2[3]};
        return IsPointInRectangel(leftBottom, rec1) || IsPointInRectangel(leftTop, rec1) || IsPointInRectangel(rightBottom, rec1) || IsPointInRectangel(rightTop, rec1);
    }

    public bool IsPointInRectangel(int[] point, int[] rec)
    {
        return (point[0] - rec[0]) * (point[0] - rec[2]) < 0 && (point[1] - rec[1]) * (point[1] - rec[3]) < 0;
    }
}


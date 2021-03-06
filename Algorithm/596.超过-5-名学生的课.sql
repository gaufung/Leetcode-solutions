--
-- @lc app=leetcode.cn id=596 lang=mysql
--
-- [596] 超过5名学生的课
--
-- https://leetcode-cn.com/problems/classes-more-than-5-students/description/
--
-- database
-- Easy (35.92%)
-- Likes:    92
-- Dislikes: 0
-- Total Accepted:    13.1K
-- Total Submissions: 36.4K
-- Testcase Example:  '{"headers": {"courses": ["student", "class"]}, "rows": {"courses": [["A", "Math"], ["B", "English"], ["C", "Math"], ["D", "Biology"], ["E", "Math"], ["F", "Computer"], ["G", "Math"], ["H", "Math"], ["I", "Math"]]}}'
--
-- 有一个courses 表 ，有: student (学生) 和 class (课程)。
-- 
-- 请列出所有超过或等于5名学生的课。
-- 
-- 例如,表:
-- 
-- 
-- +---------+------------+
-- | student | class      |
-- +---------+------------+
-- | A       | Math       |
-- | B       | English    |
-- | C       | Math       |
-- | D       | Biology    |
-- | E       | Math       |
-- | F       | Computer   |
-- | G       | Math       |
-- | H       | Math       |
-- | I       | Math       |
-- +---------+------------+
-- 
-- 
-- 应该输出:
-- 
-- 
-- +---------+
-- | class   |
-- +---------+
-- | Math    |
-- +---------+
-- 
-- 
-- Note:
-- 学生在每个课中不应被重复计算。
-- 
--
# Write your MySQL query statement below

select t.class from courses as t group by t.class having count(distinct(t.student)) >= 5

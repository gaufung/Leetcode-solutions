--
-- @lc app=leetcode.cn id=197 lang=mysql
--
-- [197] 上升的温度
--
-- https://leetcode-cn.com/problems/rising-temperature/description/
--
-- database
-- Easy (43.70%)
-- Likes:    68
-- Dislikes: 0
-- Total Accepted:    12K
-- Total Submissions: 27.4K
-- Testcase Example:  '{"headers": {"Weather": ["Id", "RecordDate", "Temperature"]}, "rows": {"Weather": [[1, "2015-01-01", 10], [2, "2015-01-02", 25], [3, "2015-01-03", 20], [4, "2015-01-04", 30]]}}'
--
-- 给定一个 Weather 表，编写一个 SQL 查询，来查找与之前（昨天的）日期相比温度更高的所有日期的 Id。
-- 
-- +---------+------------------+------------------+
-- | Id(INT) | RecordDate(DATE) | Temperature(INT) |
-- +---------+------------------+------------------+
-- |       1 |       2015-01-01 |               10 |
-- |       2 |       2015-01-02 |               25 |
-- |       3 |       2015-01-03 |               20 |
-- |       4 |       2015-01-04 |               30 |
-- +---------+------------------+------------------+
-- 
-- 例如，根据上述给定的 Weather 表格，返回如下 Id:
-- 
-- +----+
-- | Id |
-- +----+
-- |  2 |
-- |  4 |
-- +----+
-- 
--
# Write your MySQL query statement below
select t1.Id from Weather t1 
    join Weather t2 on To_Days(t1.RecordDate) = To_Days(t2.RecordDate) + 1 
        and t1.Temperature > t2.Temperature


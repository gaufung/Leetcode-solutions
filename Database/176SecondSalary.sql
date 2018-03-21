-- limit usage:
-- limit i, j: show rows from i to j (zero base)
-- limit i, -1: show rows from i to last
-- limit i: show rows from 0 to i

-- IFNULL function
-- IFNULL(expr1, expr2), if expr1 is not null return expr1 value, or else expr2

select IFNULL( 
        (select e.Salary from Employee e group by e.Salary 
            order by e.Salary desc limit 1, 1), 
         NULL
        ) 
        as SecondHighestSalary;


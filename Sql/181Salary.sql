select t1.Name as Employee 
    from Employee t1 
        inner join Employee t2 
            on t1.ManagerId = t2.Id and t1.Salary > t2.Salary
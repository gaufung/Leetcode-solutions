select t1.Id from Weather t1 
    join Weather t2 on To_Days(t1.Date) = To_Days(t2.Date) + 1 
        and t1.Temperature > t2.Temperature
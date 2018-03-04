select * from cinema t 
    where t.description != 'boring' and t.id % 2 != 0 
        order by t.rating desc
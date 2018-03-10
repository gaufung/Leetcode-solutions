delete from Person 
    where Id not in 
    (
        select temp.iid from 
        (
          (select min(Id) as iid from Person group by Email )
        ) temp
    )
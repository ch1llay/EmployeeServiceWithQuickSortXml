insert into employees
(name,
 age,
 department_name)
values (
        $1, $2, $3
       ) RETURNING id
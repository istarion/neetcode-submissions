-- Write your query below
select e.employee_id, 
    case
        when employee_id % 2 = 1 AND e.name NOT LIKE 'M%' THEN e.salary
        ELSE 0
    END as bonus
    from employees e
    order by employee_id
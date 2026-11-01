-- Write your query below
select c.customer_id, c.customer_name from customers c 
WHERe c.customer_id IN (select customer_id from orders where product_name = 'A') 
and c.customer_id IN (select customer_id from orders where product_name = 'B')
and c.customer_id NOT IN (select customer_id from orders where product_name = 'C')
order by c.customer_name
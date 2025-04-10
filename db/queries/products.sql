-- name: InsertProduct :exec
insert into products
(name, value)
values
($1, $2);

-- name: GetAllProducts :many
select * from products;

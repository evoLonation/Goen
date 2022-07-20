
-- name: CreateSale :one
insert into sales(
time, is_complete, amout, is_ready_to_pay
) values ($1, $2, $3, $4)
returning *;

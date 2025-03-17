-- name: CreateSet :one
-- insert new exercise sets
insert into 
gowebapp.sets
(Exercise_Id, Weight)
values ($1, $2)
returning *;

-- name: UpdateSet :one
-- insert a set id
update gowebapp.sets
set (Exercise_Id, Weight) = ($1, $2)
where set_id = $3
returning *;

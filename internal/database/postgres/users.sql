-- name: ListUsers :many
-- get all users ordered by username
Select * 
from 
gowebapp.users
order by user_name;

-- name: GetUser :one
-- get users of particular user_id
Select *
from
gowebapp.users
where user_id = $1;

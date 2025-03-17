-- name: ListUsers :many
-- get all users ordered by username
select * 
from 
gowebapp.users
order by user_name;

-- name: GetUser :one
-- get users of particular user_id
select *
from
gowebapp.users
where user_id = $1;

-- name: GetUserWorkout :many
-- get a particular information, exercise sets and workouts
select u.user_id, w.workout_id, w.start_date, s.set_id, s.weight
from
gowebapp.users as u,
gowebapp.workouts as w,
gowebapp.sets as s
where u.user_id = w.user_id
and w.set_id = s.set_id
and u.user_id = $1;

-- name: GetUserImage :one
-- get a particular user image
select u.name, u.user_id, i.image_data
from
gowebapp.users as u,
gowebapp.images as i
where u.user_id = i.user_id
and u.user_id = $1;

-- name: DeleteUsers :exec
-- delete a particular user
delete
from gowebapp.users
where user_id = $1;

-- name: DeleteUserImage :exec
-- delete a particular user's image
delete
from gowebapp.images as i
where i.user_id = $1;

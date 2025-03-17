-- name: CreateExercise :one
-- insert a new exercise
insert into
gowebapp.exercises (Exercise_Name)
values ($1) returning Exercise_ID;

-- name: DeleteExercise :exec
-- delete a particular exercise
delete 
from gowebapp.exercises as e
WHERE e.exercise_id = $1;

-- name: UpsertExercise :one
-- insert or update exercise of a particular id
insert into
gowebapp.exercises (Exercise_Name)
values ($1) on conflict (Exercise_ID) do
update
set Exercise_Name = Excluded.Exercise_Name
returning Exercise_ID;

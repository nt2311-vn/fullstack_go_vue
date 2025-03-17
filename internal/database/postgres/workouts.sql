-- name: CreateWorkout :one
-- insert new workouts
insert into gowebapp.workouts
(User_ID, Set_ID, Start_Date)
values ($1, $2, $3)
returning *;

-- name: UpsertWorkout :one
-- insert or update a workout based on id
insert into gowebapp.workouts
(User_ID, Set_ID, Start_Date)
values ($1, $2, $3)
on conflict (Workout_ID) do
update set
User_ID = excluded.User_ID,
Set_ID = excluded.Set_ID,
Start_Date = excluded.Start_Date
returning Workout_ID;

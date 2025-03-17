// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package pg_database

import (
	"context"
	"time"
)

const createUsers = `-- name: CreateUsers :one
insert into gowebapp.users
(User_Name, Pass_Word_Hash, name)
values ($1, $2, $3)
returning user_id, user_name, pass_word_hash, name, config, created_at, is_enabled
`

type CreateUsersParams struct {
	UserName     string
	PassWordHash string
	Name         string
}

// insert new user
func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (GowebappUser, error) {
	row := q.db.QueryRowContext(ctx, createUsers, arg.UserName, arg.PassWordHash, arg.Name)
	var i GowebappUser
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.PassWordHash,
		&i.Name,
		&i.Config,
		&i.CreatedAt,
		&i.IsEnabled,
	)
	return i, err
}

const deleteUserImage = `-- name: DeleteUserImage :exec
delete
from gowebapp.images as i
where i.user_id = $1
`

// delete a particular user's image
func (q *Queries) DeleteUserImage(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUserImage, userID)
	return err
}

const deleteUserWorkouts = `-- name: DeleteUserWorkouts :exec
delete
from gowebapp.workouts as w
where w.user_id = $1
`

// delete a particular user's workouts
func (q *Queries) DeleteUserWorkouts(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUserWorkouts, userID)
	return err
}

const deleteUsers = `-- name: DeleteUsers :exec
delete
from gowebapp.users
where user_id = $1
`

// delete a particular user
func (q *Queries) DeleteUsers(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUsers, userID)
	return err
}

const getUser = `-- name: GetUser :one
select user_id, user_name, pass_word_hash, name, config, created_at, is_enabled
from
gowebapp.users
where user_id = $1
`

// get users of particular user_id
func (q *Queries) GetUser(ctx context.Context, userID int64) (GowebappUser, error) {
	row := q.db.QueryRowContext(ctx, getUser, userID)
	var i GowebappUser
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.PassWordHash,
		&i.Name,
		&i.Config,
		&i.CreatedAt,
		&i.IsEnabled,
	)
	return i, err
}

const getUserImage = `-- name: GetUserImage :one
select u.name, u.user_id, i.image_data
from
gowebapp.users as u,
gowebapp.images as i
where u.user_id = i.user_id
and u.user_id = $1
`

type GetUserImageRow struct {
	Name      string
	UserID    int64
	ImageData []byte
}

// get a particular user image
func (q *Queries) GetUserImage(ctx context.Context, userID int64) (GetUserImageRow, error) {
	row := q.db.QueryRowContext(ctx, getUserImage, userID)
	var i GetUserImageRow
	err := row.Scan(&i.Name, &i.UserID, &i.ImageData)
	return i, err
}

const getUserWorkout = `-- name: GetUserWorkout :many
select u.user_id, w.workout_id, w.start_date, s.set_id, s.weight
from
gowebapp.users as u,
gowebapp.workouts as w,
gowebapp.sets as s
where u.user_id = w.user_id
and w.set_id = s.set_id
and u.user_id = $1
`

type GetUserWorkoutRow struct {
	UserID    int64
	WorkoutID int64
	StartDate time.Time
	SetID     int64
	Weight    int32
}

// get a particular information, exercise sets and workouts
func (q *Queries) GetUserWorkout(ctx context.Context, userID int64) ([]GetUserWorkoutRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserWorkout, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserWorkoutRow
	for rows.Next() {
		var i GetUserWorkoutRow
		if err := rows.Scan(
			&i.UserID,
			&i.WorkoutID,
			&i.StartDate,
			&i.SetID,
			&i.Weight,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
select user_id, user_name, pass_word_hash, name, config, created_at, is_enabled 
from 
gowebapp.users
order by user_name
`

// get all users ordered by username
func (q *Queries) ListUsers(ctx context.Context) ([]GowebappUser, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GowebappUser
	for rows.Next() {
		var i GowebappUser
		if err := rows.Scan(
			&i.UserID,
			&i.UserName,
			&i.PassWordHash,
			&i.Name,
			&i.Config,
			&i.CreatedAt,
			&i.IsEnabled,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

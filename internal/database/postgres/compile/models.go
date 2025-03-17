// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package pg_database

import (
	"encoding/json"
	"time"
)

type GowebappExercise struct {
	ExerciseID   int64
	ExerciseName string
}

type GowebappImage struct {
	ImageID     int64
	UserID      int64
	ContentType string
	ImageData   []byte
}

type GowebappSet struct {
	SetID      int64
	ExerciseID int64
	Weight     int32
}

type GowebappUser struct {
	UserID       int64
	UserName     string
	PassWordHash string
	Name         string
	Config       json.RawMessage
	CreatedAt    time.Time
	IsEnabled    bool
}

type GowebappWorkout struct {
	WorkoutID  int64
	SetID      int64
	UserID     int64
	ExerciseID int64
	StartDate  time.Time
}

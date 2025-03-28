// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: images.sql

package pg_database

import (
	"context"
)

const createUserImage = `-- name: CreateUserImage :one
insert into
gowebapp.images
(User_ID, Content_Type, Image_Data)
values ($1, $2, $3)
returning image_id, user_id, content_type, image_data
`

type CreateUserImageParams struct {
	UserID      int64
	ContentType string
	ImageData   []byte
}

// insert a new image
func (q *Queries) CreateUserImage(ctx context.Context, arg CreateUserImageParams) (GowebappImage, error) {
	row := q.db.QueryRowContext(ctx, createUserImage, arg.UserID, arg.ContentType, arg.ImageData)
	var i GowebappImage
	err := row.Scan(
		&i.ImageID,
		&i.UserID,
		&i.ContentType,
		&i.ImageData,
	)
	return i, err
}

const upsertUserImage = `-- name: UpsertUserImage :one
insert into
gowebapp.images
(Image_Data)
values ($1) on conflict (Image_ID) do
update
set Image_Data = excluded.Image_Data
returning Image_ID
`

// insert or update a particular image
func (q *Queries) UpsertUserImage(ctx context.Context, imageData []byte) (int64, error) {
	row := q.db.QueryRowContext(ctx, upsertUserImage, imageData)
	var image_id int64
	err := row.Scan(&image_id)
	return image_id, err
}

-- name: CreateUserImage :one
-- insert a new image
insert into
gowebapp.images
(User_ID, Content_Type, Image_Data)
values ($1, $2, $3)
returning *;

-- name: UpsertUserImage :one
-- insert or update a particular image
insert into
gowebapp.images
(Image_Data)
values ($1) on conflict (Image_ID) do
update
set Image_Data = excluded.Image_Data
returning Image_ID;

create schema if not exists gowebapp;

create table gowebapp.users (
  User_ID bigserial not null,
  User_Name text not null,
  Pass_Word_Hash text not null,
  Name text not null
  Config jsonb not null default '{}'::JSONB,
  Created_At timestamp not null default NOW(),
  Is_Enabled boolean not null default TRUE,
  Constraint PK_users primary key ( User_ID )
);

create table gowebapp.exercises (
  Exercise_ID bigserial not null,
  Exercise_Name text not null,
  Constraint PK_exercises primary key ( Exercise_ID )
);

create table gowebapp.images (
  Image_ID bigserial not null,
  User_ID bigserial not null,
  Content_Type text not null default 'image/png',
  Image_Data bytea not null,
  Constraint PK_images primary key (Image_ID, User_ID),
  Constraint FK_65 foreign key (User_ID) references gowebapp.users (User_ID)
);

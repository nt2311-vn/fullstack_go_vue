-- ************************** gowebapp
create schema if not exists gowebapp;

-- ************************** gowebapp.users
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

create index FK_67 on gowebapp.images (
  User_ID
);

create table gowebapp.sets (
  Set_ID bigserial not null,
  Exercise_ID bigserial not null,
  Weight int not null default 0,
  Constraint PK_sets primary key ( Set_ID, Exercise_ID),
  Constraint FK_106 foreign key (Exercise_ID) references gowebapp.exercises ( Exercise_ID)
);

create index FK_108 on gowebapp.sets (
  Exercise_ID
);

create table gowebapp.workouts (
  Workout_ID bigserial not null,
  Set_ID bigserial not null,
  User_ID bigserial not null,
  Exercise_ID bigserial not null,
  Start_Date timestamp not null default NOW(),
  Constraint PK_workouts primary key (Workout_ID, Set_ID, User_ID, Exercise_ID),
  Constraint FK_71 foreign key (Set_ID, Exercise_ID) references gowebapp.sets ( Set_ID, Exercise_ID)
);

create index FK_73 on gowebapp.workouts (
  Set_ID,
  Exercise_ID
);

create index FK_76 on gowebapp.workouts (
  User_ID
);

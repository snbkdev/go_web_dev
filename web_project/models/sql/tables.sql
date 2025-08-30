create table posts (
id serial primary key,
title text not null,
markdown text not null
);

create table users (
id serial primary key,
email text unique not null,
display_name text not null
);

create table comments (
id serial primary key,
user_id int,
post_id int,
markdown text not null
);
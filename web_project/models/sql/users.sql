create table users (
    id serial primary key,
    email text unique not null,
    password_hash text not null
);
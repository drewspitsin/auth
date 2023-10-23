-- +goose Up
create table user_api (
    id serial primary key,
    username text not null,
    email text not null,
    password text not null,
    role text,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table user_api;

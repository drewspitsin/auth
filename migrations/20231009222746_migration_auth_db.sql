-- +goose Up
create table testuser (
    id serial primary key,
    user_name text not null,
    email text not null,
    pswd text not null,
    role text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table user;

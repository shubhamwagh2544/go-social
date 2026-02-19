create extension if not exists citext;

create table
    if not exists users (
        id bigserial primary key,
        email citext not null unique,
        username text not null unique,
        password bytea not null,
        created_at timestamptz not null default now ()
    );
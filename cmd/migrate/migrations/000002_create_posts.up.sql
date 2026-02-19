create table
    if not exists posts (
        id bigserial primary key,
        title text not null,
        content text not null,
        user_id bigint not null,
        created_at timestamptz not null default now (),
        updated_at timestamptz not null default now ()
    );
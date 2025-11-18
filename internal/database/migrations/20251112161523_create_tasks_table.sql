-- +goose Up
-- +goose StatementBegin
create table if not exists tasks (
   id         serial primary key,
   name       varchar(255) unique not null,
   done       boolean not null default false,
   created_at timestamp with time zone not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists tasks;
-- +goose StatementEnd
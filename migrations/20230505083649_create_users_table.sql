-- +goose Up
-- с названием "user" почему-то не создавалась таблица
create table users
(
    id         bigserial primary key,
--  role_id    bigserial foreign key,
    role_id    integer   not null,
    username   text      not null,
    email      text      not null,
    password   text    not null,
    password_confirm text  not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

-- +goose Down
drop table users;
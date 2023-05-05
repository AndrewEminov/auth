-- +goose Up
create table role
(
    id      bigserial primary key,
    role    text not null
);

-- +goose Down
drop table role;

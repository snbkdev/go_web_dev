-- +goose Up
-- +goose StatementBegin
create table galleries (
    id serial primary key,
    user_id int references users (id) on delete cascade,
    title text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table galleries;
-- +goose StatementEnd

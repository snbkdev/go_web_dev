-- +goose Up
-- +goose StatementBegin
create table password_resets (
    id serial primary key,
    user_id int unique references users (id) on delete cascade,
    token_hash text unique not null,
    expires_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table password_resets;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users 
(
    id          uuid        not null primary key,
    fullname    varchar     not null,
    email       varchar     not null,
    password    varchar     not null,
    role        varchar     not null,
    created_at  timestamp   not null,
    updated_at  timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users
-- +goose StatementEnd

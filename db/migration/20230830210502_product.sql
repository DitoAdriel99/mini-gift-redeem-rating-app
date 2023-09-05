-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products 
(
    id          uuid        not null primary key,
    title       varchar     not null,
    description varchar     not null,
    points      int,
    quantity    int,
    image       varchar     not null,
    type        varchar     not null,
    banner      varchar,
    created_at  timestamp   not null,
    updated_at  timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products
-- +goose StatementEnd

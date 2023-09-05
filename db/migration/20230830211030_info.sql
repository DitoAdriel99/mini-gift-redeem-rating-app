-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product_info
(
    id          uuid        not null primary key,
    product_id  uuid        not null,
    info        varchar     not null,
    created_at  timestamp   not null,
    updated_at  timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product_info
-- +goose StatementEnd

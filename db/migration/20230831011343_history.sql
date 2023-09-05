-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product_history
(
    id          uuid        not null primary key,
    product_id  uuid        not null,
    email_user  varchar     not null,
    quantity    int,
    rating      float,
    created_at  timestamp   not null,
    updated_at  timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product_history
-- +goose StatementEnd

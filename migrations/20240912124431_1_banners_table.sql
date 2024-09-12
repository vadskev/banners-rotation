-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS banners_table
(
    id         SERIAL   CONSTRAINT banners_pk PRIMARY KEY,
    name       VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS banners;
-- +goose StatementEnd

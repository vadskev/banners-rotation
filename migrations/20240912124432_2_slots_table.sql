-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS slots_table
(
    id         SERIAL   CONSTRAINT slots_pk PRIMARY KEY,
    name       VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS slots_table;
-- +goose StatementEnd

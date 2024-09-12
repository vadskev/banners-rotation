-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS social_group_table
(
    id         SERIAL   CONSTRAINT social_group_pk PRIMARY KEY,
    name       VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS social_group_table;
-- +goose StatementEnd

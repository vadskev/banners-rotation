-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS banner_slot_table
(
    banner_id serial NOT NULL,
    slot_id   serial NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (banner_id, slot_id),
    FOREIGN KEY (banner_id) REFERENCES banners_table (id),
    FOREIGN KEY (slot_id) REFERENCES slots_table (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS banner_slot_table;
-- +goose StatementEnd

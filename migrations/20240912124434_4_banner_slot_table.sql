-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS banner_slot_table
(
    banner_id INT NOT NULL CONSTRAINT banner_slot_banners_id_fk REFERENCES banners_table(id) ON UPDATE CASCADE ON DELETE CASCADE,
    slot_id INT NOT NULL CONSTRAINT banner_slot_slots_id_fk REFERENCES slots_table(id) ON UPDATE CASCADE ON DELETE CASCADE,
    created_at NOT NULL DEFAULT current_timestamp,
    CONSTRAINT banner_slot_pk PRIMARY KEY (banner_id, slot_id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS banner_slot_table;
-- +goose StatementEnd

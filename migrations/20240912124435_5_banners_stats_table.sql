-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS banners_stats_table
(
    banner_id INT REFERENCES banners_table (id) ON DELETE SET NULL,
    slot_id INT REFERENCES slots_table (id) ON DELETE SET NULL,
    social_group_id INT REFERENCES social_group_table (id) ON DELETE SET NULL,
    clicks_amount int NOT NULL,
    shows_amount int NOT NULL,
    CONSTRAINT banners_stats_pk PRIMARY KEY (slot_id, social_group_id, banner_id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS banners_stats_table;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE deleted_member_travel (
    id int NOT NULL,
    member_id int unsigned NOT NULL,
    travel_id int unsigned NOT NULL,
    created_at datetime NOT NULL,
    deleted_at datetime NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE deleted_member_travel;
-- +goose StatementEnd

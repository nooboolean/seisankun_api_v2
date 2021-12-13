-- +goose Up
-- +goose StatementBegin
CREATE TABLE member_travel (
    id int unsigned NOT NULL AUTO_INCREMENT,
    member_id int unsigned NOT NULL,
    travel_id int unsigned NOT NULL,
    created_at datetime NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY fk_member_id(member_id) REFERENCES members(id),
    FOREIGN KEY fk_travel_id(travel_id) REFERENCES travels(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE member_travel;
-- +goose StatementEnd

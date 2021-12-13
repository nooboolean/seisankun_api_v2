-- +goose Up
-- +goose StatementBegin
CREATE TABLE travels (
    id int unsigned NOT NULL AUTO_INCREMENT,
    name varchar(50) NOT NULL,
    travel_key varchar(255) NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE travels;
-- +goose StatementEnd

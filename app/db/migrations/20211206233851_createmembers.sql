-- +goose Up
-- +goose StatementBegin
CREATE TABLE members (
    id int unsigned NOT NULL AUTO_INCREMENT,
    name varchar(20) NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE members;
-- +goose StatementEnd

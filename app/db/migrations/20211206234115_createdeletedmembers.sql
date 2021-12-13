-- +goose Up
-- +goose StatementBegin
CREATE TABLE deleted_members (
    id int NOT NULL,
    name varchar(20) NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE deleted_members;
-- +goose StatementEnd

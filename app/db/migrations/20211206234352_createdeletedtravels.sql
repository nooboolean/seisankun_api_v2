-- +goose Up
-- +goose StatementBegin
CREATE TABLE deleted_travels (
    id int NOT NULL,
    name varchar(50) NOT NULL,
    travel_key varchar(255) NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE deleted_travels;
-- +goose StatementEnd

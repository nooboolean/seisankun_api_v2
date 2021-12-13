-- +goose Up
-- +goose StatementBegin
CREATE TABLE deleted_payments (
    id int NOT NULL,
    travel_id int NOT NULL,
    payer_id int NOT NULL,
    title varchar(30) NOT NULL,
    amount int NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE deleted_payments;
-- +goose StatementEnd

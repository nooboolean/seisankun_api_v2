-- +goose Up
-- +goose StatementBegin
CREATE TABLE payments (
    id int unsigned NOT NULL AUTO_INCREMENT,
    travel_id int unsigned NOT NULL,
    payer_id int unsigned NOT NULL,
    title varchar(30) NOT NULL,
    amount int NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY fk_travel_id(travel_id) REFERENCES travels(id),
    FOREIGN KEY fk_payer_id(payer_id) REFERENCES members(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payments;
-- +goose StatementEnd

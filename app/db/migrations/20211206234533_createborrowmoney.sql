-- +goose Up
-- +goose StatementBegin
CREATE TABLE borrow_money (
    id int unsigned NOT NULL AUTO_INCREMENT,
    payment_id int unsigned NOT NULL,
    borrower_id int unsigned NOT NULL,
    money double NOT NULL DEFAULT '0',
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY fk_payment_id(payment_id) REFERENCES payments(id),
    FOREIGN KEY fk_borrower_id(borrower_id) REFERENCES members(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE borrow_money;
-- +goose StatementEnd

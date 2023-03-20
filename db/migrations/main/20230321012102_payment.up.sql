BEGIN;

CREATE TABLE IF NOT EXISTS main.payments
(
    "id"            BIGINT          NOT NULL PRIMARY KEY,
    "amount"        BIGINT          NOT NULL,
    "order_id"      BIGINT          NOT NULL,
    "order_status"  VARCHAR         NOT NULL DEFAULT ''
);
INSERT INTO main.payments (id, amount, order_id, order_status) VALUES(1, 90000, 123, 'Paid');
INSERT INTO main.payments (id, amount, order_id, order_status) VALUES(2, 120000, 124, 'Unpaid');
INSERT INTO main.payments (id, amount, order_id, order_status) VALUES(3, 50000, 125, 'Paid');

COMMIT;

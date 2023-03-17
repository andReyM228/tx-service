-- +migrate Up
CREATE TABLE IF NOT EXISTS balances(
                                    id BIGSERIAL PRIMARY KEY NOT NULL,
                                    user_id NUMERIC NOT NULL,
                                    amount NUMERIC NOT NULL,
                                    created_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS transactions(
                                       id BIGSERIAL PRIMARY KEY NOT NULL,
                                       user_id_from NUMERIC NOT NULL,
                                       user_id_to NUMERIC NOT NULL,
                                       amount NUMERIC NOT NULL,
                                       created_at TIMESTAMP DEFAULT current_timestamp
);

-- +migrate Down

DROP TABLE IF EXISTS balances;
DROP TABLE IF EXISTS transactions;
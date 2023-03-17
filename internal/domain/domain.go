package domain

import "time"

type Balances struct {
	ID        int
	UserID    int `db:"user_id"`
	Amount    int64
	CreatedAt time.Time `db:"created_at"`
}

type Transactions struct {
	ID         int
	UserIDFrom int `db:"user_id_from"`
	UserIDTo   int `db:"user_id_to"`
	Amount     int64
	CreatedAt  time.Time `db:"created_at"`
}

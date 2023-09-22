package transactions

import (
	"database/sql"
	"errors"
	"github.com/andReyM228/lib/log"

	"tx_service/internal/domain"
	"tx_service/internal/repository"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db  *sqlx.DB
	log log.Logger
}

func NewRepository(database *sqlx.DB, log log.Logger) Repository {
	return Repository{
		db:  database,
		log: log,
	}
}

func (r Repository) Get(id int64) (domain.Transactions, error) {
	var transaction domain.Transactions

	if err := r.db.Get(&transaction, "SELECT * FROM transactions WHERE id = $1", id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Error(err.Error())
			return domain.Transactions{}, repository.NotFound{NotFound: "transaction"}
		}

		r.log.Error(err.Error())
		return domain.Transactions{}, repository.InternalServerError{}
	}

	return transaction, nil
}

func (r Repository) Update(transaction domain.Transactions) error {
	_, err := r.db.Exec("UPDATE transactions SET user_id_from = $1, user_id_to = $2, amount = $3 WHERE id = $4", transaction.UserIDFrom, transaction.UserIDTo, transaction.Amount, transaction.ID)

	if err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Create(transaction domain.Transactions) error {
	if _, err := r.db.Exec("INSERT INTO transactions (user_id_from, user_id_to, amount) VALUES ($1, $2, $3)", transaction.UserIDFrom, transaction.UserIDTo, transaction.Amount); err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM transactions WHERE id = $1", id)
	if err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}

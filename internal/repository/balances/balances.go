package balances

import (
	"database/sql"
	"errors"

	"tx_service/internal/domain"
	"tx_service/internal/repository"

	"github.com/andReyM228/lib/log"
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

func (r Repository) Get(userID int64) (domain.Balances, error) {
	var balance domain.Balances

	if err := r.db.Get(&balance, "SELECT * FROM balances WHERE user_id = $1", userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Info(err.Error())
			return domain.Balances{}, repository.NotFound{NotFound: "balance"}
		}

		r.log.Error(err.Error())
		return domain.Balances{}, repository.InternalServerError{}
	}

	return balance, nil
}

func (r Repository) Update(balance domain.Balances) error {
	_, err := r.db.Exec("UPDATE balances SET user_id = $1, amount = $2 WHERE id = $3", balance.UserID, balance.Amount, balance.ID)

	if err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Create(balance domain.Balances) error {
	if _, err := r.db.Exec("INSERT INTO balances (user_id, amount) VALUES ($1, $2)", balance.UserID, balance.Amount); err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM balances WHERE id = $1", id)
	if err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}

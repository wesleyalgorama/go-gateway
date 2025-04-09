package repository

import (
	"database/sql"

	"github.com/wesleyalgorama/fcw/go-gateway/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	smtp, err := r.db.Prepare(`
		INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)

	if err != nil {
		return err
	}

	defer smtp.Close()

	_, err = smtp.Exec(
		account.Id,
		account.Name,
		account.Email,
		account.ApiKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

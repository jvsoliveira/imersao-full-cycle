package factory

import (
	"database/sql"

	repo "github.com/jvsoliveira/imersao-full-cycle/adapter/repository"
	"github.com/jvsoliveira/imersao-full-cycle/domain/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repo.NewTransactionRepositoryDb(r.DB)
}

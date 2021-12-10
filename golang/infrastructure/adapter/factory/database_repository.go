package factory

import (
	"database/sql"
	"github.com/loxt/imersao-fullstack-fullcycle-5/domain/repository"
	repository2 "github.com/loxt/imersao-fullstack-fullcycle-5/infrastructure/adapter/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repository2.NewTransactionRepositoryDb(r.DB)
}

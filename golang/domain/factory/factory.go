package factory

import "github.com/loxt/imersao-fullstack-fullcycle-5/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}

package factory

import "github.com/jvsoliveira/imersao-full-cycle/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}

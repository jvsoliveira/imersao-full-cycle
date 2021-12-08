package factory

import "github.com/jvsoliveira/imersao-full-cycle-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}

package manager

import (
	"echFundamental/infra"
	"echFundamental/useCases"
)

type UseCaseManager interface {
	CustomerUseCase() useCases.CustomerUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (ucm *useCaseManager) CustomerUseCase() useCases.CustomerUseCase {
	return useCases.NewCustomerUseCase(ucm.repoManager.CustomerRepo())
}

func NewUseCaseManager(infra infra.Infra) UseCaseManager {
	return &useCaseManager{NewRepoManager(infra)}
}

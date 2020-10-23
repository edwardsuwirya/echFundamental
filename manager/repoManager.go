package manager

import (
	"echFundamental/repositories"
)

type RepoManager interface {
	CustomerRepo() repositories.CustomerRepo
}

type repoManager struct {
}

func (rm *repoManager) CustomerRepo() repositories.CustomerRepo {
	return repositories.NewCustomerRepo()
}

func NewRepoManager() RepoManager {
	return &repoManager{}
}

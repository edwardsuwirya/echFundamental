package manager

import (
	"database/sql"
	"echFundamental/infra"
	"echFundamental/repositories"
)

type RepoManager interface {
	CustomerRepo() repositories.CustomerRepo
}

type repoManager struct {
	sqlConn *sql.DB
	infra   infra.Infra
}

func (rm *repoManager) CustomerRepo() repositories.CustomerRepo {
	return repositories.NewCustomerRepo(rm.sqlConn)
}

func NewRepoManager(infra infra.Infra) RepoManager {
	return &repoManager{infra.SqlDb(), infra}
}

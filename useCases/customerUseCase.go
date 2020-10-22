package useCases

import (
	"echFundamental/models"
	"echFundamental/repositories"
)

type CustomerUseCase interface {
	Register(cust *models.Customer) error
	UserInfo(id string) (*models.Customer, error)
	UserList(pageNo, totalPerPage int) ([]*models.Customer, error)
}

type CustomerUseCaseImpl struct {
	customerRepo repositories.CustomerRepo
}

func (c *CustomerUseCaseImpl) Register(cust *models.Customer) error {
	if err := c.customerRepo.Save(cust); err != nil {
		return err
	}
	return nil
}
func (c *CustomerUseCaseImpl) UserList(pageNo, totalPerPage int) ([]*models.Customer, error) {
	//Ini tidak boleh dilakukan, berbahaya, resiko tanggung sendiri
	userList := c.customerRepo.FindAll()
	var start = (pageNo - 1) * totalPerPage
	var end = pageNo * totalPerPage
	if start < 0 {
		start = 0
	}
	if end > len(userList) {
		end = len(userList)
	}
	userPaging := userList[start:end]
	return userPaging, nil
}
func (c *CustomerUseCaseImpl) UserInfo(id string) (*models.Customer, error) {
	//tugas
	return nil, nil
}

func NewCustomerUseCase(customerRepo repositories.CustomerRepo) CustomerUseCase {
	return &CustomerUseCaseImpl{
		customerRepo,
	}
}

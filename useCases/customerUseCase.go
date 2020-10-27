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
	var start = (pageNo - 1) * totalPerPage
	if start < 0 {
		start = 0
	}
	userList := c.customerRepo.FindAll(start, totalPerPage)
	return userList, nil
}
func (c *CustomerUseCaseImpl) UserInfo(id string) (*models.Customer, error) {
	return c.customerRepo.FindOne(id)
}

func NewCustomerUseCase(customerRepo repositories.CustomerRepo) CustomerUseCase {
	return &CustomerUseCaseImpl{
		customerRepo,
	}
}

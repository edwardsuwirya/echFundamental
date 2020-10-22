package repositories

import (
	"echFundamental/models"
	guuid "github.com/google/uuid"
)

type CustomerRepo interface {
	Save(c *models.Customer) error
	Update(id string, newCustomer *models.Customer) error
	Delete(id string) error
	FindAll() []*models.Customer
}

type CustomerRepoImpl struct {
	customerList []*models.Customer
}

func (c *CustomerRepoImpl) Save(cust *models.Customer) error {
	id := guuid.New()
	cust.SetId(id.String())
	c.customerList = append(c.customerList, cust)
	return nil
}

func (c *CustomerRepoImpl) Update(id string, newCustomer *models.Customer) error {
	panic("implement me")
}

func (c *CustomerRepoImpl) Delete(id string) error {
	panic("implement me")
}

func (c *CustomerRepoImpl) FindAll() []*models.Customer {
	return c.customerList
}
func NewCustomerRepo() CustomerRepo {
	repo := make([]*models.Customer, 0)
	return &CustomerRepoImpl{
		customerList: repo,
	}
}

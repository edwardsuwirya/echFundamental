package repositories

import (
	"database/sql"
	"echFundamental/models"
	guuid "github.com/google/uuid"
)

type CustomerRepo interface {
	Save(c *models.Customer) error
	Update(id string, newCustomer *models.Customer) error
	Delete(id string) error
	FindAll(pageNo, totalPerPage int) []*models.Customer
}

type CustomerRepoImpl struct {
	db *sql.DB
}

func (c *CustomerRepoImpl) Save(cust *models.Customer) error {
	id := guuid.New()
	cust.SetId(id.String())
	sql := "INSERT INTO M_CUSTOMER(id,first_name,last_name,address,city) VALUES(?,?,?,?,?)"
	stmt, err := c.db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cust.Id, cust.FirstName, cust.LastName, cust.Address.Address, cust.Address.City)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerRepoImpl) Update(id string, newCustomer *models.Customer) error {
	panic("implement me")
}

func (c *CustomerRepoImpl) Delete(id string) error {
	panic("implement me")
}

func (c *CustomerRepoImpl) FindAll(pageNo, totalPerPage int) []*models.Customer {
	sql := "SELECT * FROM M_CUSTOMER LIMIT ?,?"
	stmt, err := c.db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(pageNo, totalPerPage)
	if err != nil {
		panic(err)
	}
	var customerList []*models.Customer
	for rows.Next() {
		customer := new(models.Customer)
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Address.Address, &customer.Address.City)
		if err != nil {
			panic(err)
		}
		customerList = append(customerList, customer)
	}
	return customerList
}
func NewCustomerRepo(dbConn *sql.DB) CustomerRepo {
	return &CustomerRepoImpl{
		db: dbConn,
	}
}

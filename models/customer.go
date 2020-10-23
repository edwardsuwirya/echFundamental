package models

import "fmt"

type Customer struct {
	Id        string
	FirstName string
	LastName  string
	Address   CustomerAddress
}

func (c *Customer) SetId(id string) {
	c.Id = id
}
func (c *Customer) ToString() string {
	return fmt.Sprintf("Customer (id: %s) with full name %s %s, address %s",
		c.Id, c.FirstName, c.LastName, c.Address.Address)
}

func NewCustomerDefault() *Customer {
	return &Customer{}
}
func NewCustomer(firstName, lastName, address string) *Customer {
	return &Customer{
		FirstName: firstName,
		LastName:  lastName,
		Address: CustomerAddress{
			Address: address,
			City:    "",
		},
	}
}

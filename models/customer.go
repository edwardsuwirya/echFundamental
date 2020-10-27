package models

import "fmt"

type Customer struct {
	//struct tag
	Id        string          `json:"id""`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Address   CustomerAddress `json:"address"`
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

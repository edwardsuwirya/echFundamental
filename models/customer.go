package models

import "fmt"

type Customer struct {
	id        string
	firstName string
	lastName  string
	address   CustomerAddress
}

func (c *Customer) SetId(id string) {
	c.id = id
}
func (c *Customer) ToString() string {
	return fmt.Sprintf("Customer (id: %s) with full name %s %s, address %s",
		c.id, c.firstName, c.lastName, c.address.address)
}

func NewCustomerDefault() *Customer {
	return &Customer{}
}
func NewCustomer(firstName, lastName, address string) *Customer {
	return &Customer{
		firstName: firstName,
		lastName:  lastName,
		address: CustomerAddress{
			address: address,
			city:    "",
		},
	}
}

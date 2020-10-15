package models

import "fmt"


type Customer struct {
	firstName string
	lastName string
	address CustomerAddress
}

func (c *Customer) ToString() string{
	return fmt.Sprintf("Customer with full name %s %s, address %s",
		c.firstName,c.lastName,c.address.address)
}

func NewCustomerDefault() *Customer{
	return &Customer{}
}
func NewCustomer(firstName,lastName, address string) *Customer{
	return &Customer{
		firstName: firstName,
		lastName:  lastName,
		address:   CustomerAddress{
			address: address,
			city:    "",
		},
	}
}

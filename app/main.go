package main

import (
	"echFundamental/models"
	"echFundamental/repositories"
	"fmt"
)

func main() {
	repo := repositories.NewCustomerRepo()
	customer1 := models.NewCustomer("Arip","Komeng","Ragunan")
	repo.Save(customer1)
	customer2 := models.NewCustomer("Ka","Edi","Ragunan")
	repo.Save(customer2)
	customerList := repo.FindAll()

	for _,c:=range customerList{
		fmt.Println(c.ToString())
	}
}


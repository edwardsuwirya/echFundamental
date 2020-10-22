package deliveries

import (
	"echFundamental/models"
	"echFundamental/useCases"
	"fmt"
	"log"
)

type CustomerDelivery struct {
	customerUseCase useCases.CustomerUseCase
}

func NewCustomerDelivery(customerUseCase useCases.CustomerUseCase) *CustomerDelivery {
	return &CustomerDelivery{customerUseCase}
}

func (cd *CustomerDelivery) RegisterCustomerBulk(customerBulk ...*models.Customer) {
	for _, c := range customerBulk {
		if err := cd.customerUseCase.Register(c); err != nil {
			panic(err)
		}
	}
}

func (cd *CustomerDelivery) PrintConsoleCustomerPaging(pageNo, totalPerPage int) {
	userList, err := cd.customerUseCase.UserList(pageNo, totalPerPage)
	if err != nil {
		log.Println(err)
	}
	for _, c := range userList {
		fmt.Println(c.ToString())
	}
}

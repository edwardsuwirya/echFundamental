package main

import (
	"echFundamental/deliveries"
	"echFundamental/models"
	"echFundamental/repositories"
	"echFundamental/useCases"
)

type App struct{}

func (app *App) run() {
	customerRepo := repositories.NewCustomerRepo()
	customerUseCase := useCases.NewCustomerUseCase(customerRepo)
	customerDelivery := deliveries.NewCustomerDelivery(customerUseCase)

	customer1 := models.NewCustomer("Arip", "Komeng", "Ragunan")
	customer2 := models.NewCustomer("Ka", "Edi", "Ragunan")
	customer3 := models.NewCustomer("Aldy", "Boii", "Ragunan")
	customer4 := models.NewCustomer("Egi", "Mahmud", "Pasar Minggu")
	customer5 := models.NewCustomer("Winda", "Ayu", "Fatmawati")

	customerDelivery.RegisterCustomerBulk(customer1, customer2, customer3, customer4, customer5)
	customerDelivery.PrintConsoleCustomerPaging(1, 2)
}

func NewApp() *App {
	return &App{}
}
func main() {
	NewApp().run()
}

package main

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func (app *App) run() {
	//customerRepo := repositories.NewCustomerRepo()
	//customerUseCase := useCases.NewCustomerUseCase(customerRepo)
	//customerDelivery := deliveries.NewCustomerDelivery(customerUseCase)
	//
	//customer1 := models.NewCustomer("Arip", "Komeng", "Ragunan")
	//customer2 := models.NewCustomer("Ka", "Edi", "Ragunan")
	//customer3 := models.NewCustomer("Aldy", "Boii", "Ragunan")
	//customer4 := models.NewCustomer("Egi", "Mahmud", "Pasar Minggu")
	//customer5 := models.NewCustomer("Winda", "Ayu", "Fatmawati")
	//
	//customerDelivery.RegisterCustomerBulk(customer1, customer2, customer3, customer4, customer5)
	//customerDelivery.PrintConsoleCustomerPaging(1, 2)
	app.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	if err := app.router.Run("localhost:6969"); err != nil {
		panic(err)
	}

}

func NewApp() *App {
	//https://github.com/gin-gonic/gin
	router := gin.Default()
	return &App{router}
}
func main() {
	NewApp().run()
}

package main

import (
	"echFundamental/deliveries"
	"echFundamental/infra"
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
	infra  infra.Infra
}

func (app *App) run() {
	deliveries.NewAppRouter(app.infra, app.router).InitRouter()
	if err := app.router.Run(app.infra.ApiServer()); err != nil {
		panic(err)
	}
}

func NewApp() *App {
	//https://github.com/gin-gonic/gin
	router := gin.Default()
	appInfra := infra.NewInfra()
	return &App{router, appInfra}
}
func main() {
	NewApp().run()
}

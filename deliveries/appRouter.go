package deliveries

import (
	"echFundamental/manager"
	"github.com/gin-gonic/gin"
)

type appRouter struct {
	router *gin.Engine
}

func NewAppRouter(router *gin.Engine) *appRouter {
	return &appRouter{router}
}

func (ar *appRouter) InitRouter() {
	useCaseManager := manager.NewUseCaseManager()
	ar.router.GET("/ping", pingController)
	NewCustomerDelivery(ar.router, useCaseManager.CustomerUseCase()).InitRoute()
}

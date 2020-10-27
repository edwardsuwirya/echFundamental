package deliveries

import (
	"echFundamental/httpUtils/httpResponse"
	"echFundamental/infra"
	"echFundamental/manager"
	"github.com/gin-gonic/gin"
)

type appRouter struct {
	router *gin.Engine
	infra  infra.Infra
}

func NewAppRouter(infra infra.Infra, router *gin.Engine) *appRouter {
	return &appRouter{router, infra}
}

func (ar *appRouter) InitRouter() {
	useCaseManager := manager.NewUseCaseManager(ar.infra)
	ar.router.GET("/ping", pingController)
	NewCustomerDelivery(ar.router, useCaseManager.CustomerUseCase(), httpResponse.NewJsonResponder()).InitRoute()
}

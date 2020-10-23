package deliveries

import (
	"github.com/gin-gonic/gin"
)

type appRouter struct {
	router *gin.Engine
}

func NewAppRouter(router *gin.Engine) *appRouter {
	return &appRouter{router}
}

func (ar *appRouter) InitRouter() {
	ar.router.GET("/ping", pingController)
	NewCustomerDelivery(ar.router).InitRoute()
}

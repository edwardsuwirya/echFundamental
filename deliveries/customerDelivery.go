package deliveries

import (
	"echFundamental/repositories"
	"echFundamental/useCases"
	"github.com/gin-gonic/gin"
)

type CustomerDelivery struct {
	router *gin.Engine
}

func NewCustomerDelivery(router *gin.Engine) AppDelivery {
	return &CustomerDelivery{router}
}

func (cd *CustomerDelivery) InitRoute() {
	custRepo := repositories.NewCustomerRepo()
	custUseCase := useCases.NewCustomerUseCase(custRepo)
	custController := NewCustomerController(custUseCase)
	customerRouter := cd.router.Group("/customer")
	customerRouter.GET("", custController.getCustomerPagingController)
}

package deliveries

import (
	"echFundamental/httpUtils/httpResponse"
	"echFundamental/useCases"
	"github.com/gin-gonic/gin"
)

type CustomerDelivery struct {
	router     *gin.Engine
	controller *CustomerController
}

func NewCustomerDelivery(router *gin.Engine, custUseCase useCases.CustomerUseCase, responder httpResponse.IResponder) AppDelivery {
	return &CustomerDelivery{router, NewCustomerController(custUseCase, responder)}
}

func (cd *CustomerDelivery) InitRoute() {
	customerRouter := cd.router.Group("/customer")
	customerRouter.GET("/:id", cd.controller.getCustomerById)
	customerRouter.GET("", cd.controller.getCustomerPagingController)
	customerRouter.POST("", cd.controller.registerCustomer)
}

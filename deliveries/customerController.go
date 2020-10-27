package deliveries

import (
	"echFundamental/httpUtils/httpResponse"
	"echFundamental/models"
	"echFundamental/useCases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CustomerController struct {
	customerUseCase useCases.CustomerUseCase
	responder       httpResponse.IResponder
}

func NewCustomerController(useCase useCases.CustomerUseCase, responder httpResponse.IResponder) *CustomerController {
	return &CustomerController{customerUseCase: useCase, responder: responder}
}
func (cc *CustomerController) registerCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		cc.responder.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
	}
	if err := cc.customerUseCase.Register(&customer); err != nil {
		cc.responder.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
	}
	cc.responder.SetContext(c).SingleResponder(http.StatusOK, customer)
}

func (cc *CustomerController) getCustomerById(c *gin.Context) {
	custId := c.Param("id")
	customer, err := cc.customerUseCase.UserInfo(custId)
	if err != nil {
		cc.responder.SetContext(c).ErrorResponder(http.StatusBadRequest, "", err.Error())
	}
	cc.responder.SetContext(c).SingleResponder(http.StatusOK, customer)
}
func (cc *CustomerController) getCustomerPagingController(c *gin.Context) {
	// http://localhost:8080/customer?pageNo=1&totalPerPage=10
	pageNo := c.DefaultQuery("pageNo", "1")
	totalPerPage := c.DefaultQuery("totalPerPage", "5")

	iPageNo, _ := strconv.Atoi(pageNo)
	iTotalPerPage, _ := strconv.Atoi(totalPerPage)
	customerList, err := cc.customerUseCase.UserList(iPageNo, iTotalPerPage)
	if err != nil {
		cc.responder.SetContext(c).ErrorResponder(http.StatusInternalServerError, "", err.Error())
	}
	cc.responder.SetContext(c).PagingResponder(http.StatusOK, customerList, iPageNo, len(customerList))
}

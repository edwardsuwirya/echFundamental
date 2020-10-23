package deliveries

import (
	"echFundamental/useCases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CustomerController struct {
	customerUseCase useCases.CustomerUseCase
}

func NewCustomerController(useCase useCases.CustomerUseCase) *CustomerController {
	return &CustomerController{customerUseCase: useCase}
}
func (cc *CustomerController) getCustomerPagingController(c *gin.Context) {
	// http://localhost:8080/customer?pageNo=1&totalPerPage=10
	pageNo := c.DefaultQuery("pageNo", "1")
	totalPerPage := c.DefaultQuery("totalPerPage", "5")

	iPageNo, _ := strconv.Atoi(pageNo)
	iTotalPerPage, _ := strconv.Atoi(totalPerPage)
	customerList, err := cc.customerUseCase.UserList(iPageNo, iTotalPerPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": customerList,
	})
}

package deliveries

import (
	"echFundamental/models"
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
func (cc *CustomerController) registerCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := cc.customerUseCase.Register(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": customer,
	})
}

func (cc *CustomerController) getCustomerById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "123",
	})
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
		"data":      customerList,
		"pageNo":    iPageNo,
		"totalData": len(customerList),
	})
}

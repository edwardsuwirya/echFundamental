package httpResponse

type IResponder interface {
	PagingResponder(successCode int, data interface{}, pageNo int, totalData int)
	SingleResponder(successCode int, data interface{})
	ErrorResponder(errorCode int, appErrorCode string, errorMessage string)
	SetContext(ctx interface{}) IResponder
}

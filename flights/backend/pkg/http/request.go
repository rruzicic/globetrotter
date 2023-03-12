package http

import (
	"fmt"
)

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Bindes request body JSON to forwarded body interface, if there are validation errors it will respond with
// a list of ErrorMessage objects. If no error is present it will return &body object with populated fields.
func (gin *Gin) BindAndValidateBody(body interface{}) *interface{} {
	if err := gin.Context.ShouldBindJSON(&body); err != nil {
		fmt.Println(err.Error())
		/*var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMessage, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMessage{
					Field:   fe.Field(),
					Message: fe.Error(),
				}
			}
			}*/
		sendResponse(gin.Context, 400, "Bad request", "could not bind or validate data")
		//gin.Context.AbortWithStatusJSON(400, "could not bind or validate data")
	}
	return &body
}

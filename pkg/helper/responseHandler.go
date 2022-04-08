package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Count      int         `json:"count"`
}

// Empty object is used when data doenst want to be null on json
type EmptyObj struct{}

func InternalServerError() {
	res := Response{
		Success: false,
		Message: "INTERNALSERVERERROR",
		Data:    nil,
		Count:   0,
	}
	APIResponseHandler(res, 500)
}

// Used to inject data value to dynamic failed response
//status 404
func NotFoundError() {
	fmt.Println("notfound err")

	res := Response{
		Success: false,
		Message: "NOTFOUND",
		Data:    nil,
		Count:   0,
	}
	APIResponseHandler(res, 404)
}

// status 401
func Unauthorized() {
	res := Response{
		Success: false,
		Message: "UNAUTHORIZED",
		Data:    nil,
		Count:   0,
	}
	APIResponseHandler(res, 401)
}

// Used to inject data value to dynamic failed response
//status 400
func BadRequestError() {
	res := Response{
		Success: false,
		Message: "BADREQUEST",
		Data:    nil,
		Count:   0,
	}
	APIResponseHandler(res, 400)
}

// status 400
func MailPhoneError() {
	res := Response{
		Success: false,
		Message: "ERRORMAILORPHONE",
		Data:    nil,
		Count:   0,
	}
	APIResponseHandler(res, 400)
}

// status 400
func ImageError() {
	res := Response{
		Success: false,
		Message: "ERRORIMG",
		Data:    nil,
		Count:   0,
	}
	APIResponseHandler(res, 400)
}

//status 403
func ForbiddenError() {
	res := Response{
		Success: false,
		Message: "FORBIDDEN",
		Data:    nil,
		Count:   0,
	}
	APIResponseHandler(res, 403)
}

// Used to inject data value to dynamic failed response
// status 200
func ErrorLogin() {
	res := Response{
		Success: false,
		Message: "ERRORLOGIN",
		Data:    nil,
		Count:   0,
	}
	APIResponseHandler(res, 200)
}

// status 200
func WrongPassword() {
	res := Response{
		Success: false,
		Message: "WRONGPASSWORD",
		Data:    nil,
		Count:   0,
	}
	APIResponseHandler(res, 200)
}

// status 200
func Unverified() {
	res := Response{
		Success: false,
		Message: "UNVERIFIED",
		Data:    nil,
		Count:   0,
	}
	APIResponseHandler(res, 200)
}

// status 200
func Registered() {
	res := Response{
		Success: false,
		Message: "REGISTERED",
		Data:    nil,
		Count:   0,
	}
	APIResponseHandler(res, 200)
}

// status 200
func IsBanned() {
	res := Response{
		Success: false,
		Message: "ISBANNED",
		Data:    nil,
		Count:   0,
	}
	APIResponseHandler(res, 200)
}

// status 200
func OK(data interface{}) {
	res := Response{
		Success: false,
		Message: "OK",
		Data:    data,
		Count:   1,
	}
	APIResponseHandler(res, 200)
}

func APIResponseHandler(h Response, statusCode int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(statusCode, h)
	}
}

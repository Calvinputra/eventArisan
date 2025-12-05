package helper

import (
	"event/backend/model"
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Catch(activity string, c *gin.Context) {
	if r := recover(); r != nil {
		fmt.Printf("Internal error while %s, %s\r\n %s\n", activity, r, string(debug.Stack()))
		c.JSON(500, model.Response{
			HttpCode:    500,
			Message:     "ERROR_INTERNAL_SERVER_ERROR",
			MessageId:   "",
			MessageCode: 99,
			Errors:      nil,
			Data:        nil,
			NextPage:    nil,
		})
	}
}

func CatchActivity(activity string) {
	if r := recover(); r != nil {
		fmt.Printf("Internal error while %s, %s\r\n %s\n", activity, r, string(debug.Stack()))
	}
}

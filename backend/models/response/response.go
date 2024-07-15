package response

import "github.com/gin-gonic/gin"

type Response struct {
	Error     string      `json:"error"`
	ErrorCode uint        `json:"error_code"`
	Succes    bool        `json:"success"`
	Data      interface{} `json:"data"`
}

func Success(data interface{}, c *gin.Context) {
	c.JSON(200, Response{
		Succes: true,
		Data:   data,
	})
}

func Error(err string, errorCode uint, c *gin.Context) {
	c.JSON(200, Response{
		Error:     err,
		ErrorCode: errorCode,
		Succes:    false,
		Data:      nil,
	})
}

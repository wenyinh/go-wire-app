package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`    // 0: success, 非0: fail
	Message string      `json:"message"` // 描述信息
	Data    interface{} `json:"data"`    // 响应数据，可为 nil
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    nil,
	})
}

func SuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func Fail(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    1,
		Message: message,
		Data:    nil,
	})
}

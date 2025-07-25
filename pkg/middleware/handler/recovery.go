package handler

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RecoveryHandler 捕获 panic，打印日志，并返回统一错误响应
func RecoveryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				zap.L().Error("panic recovered",
					zap.Any("recover", rec),
					zap.String("stack", string(debug.Stack())),
				)

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"Code":    "InternalServerError",
					"Message": "Internal server error",
				})
			}
		}()
		c.Next()
	}
}

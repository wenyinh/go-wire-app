package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const HealthStatusUri = "/health/status"

func GetHealthStatus(c *gin.Context) {
	// HEAD method must not send message body in response
	// ref: https://datatracker.ietf.org/doc/html/rfc7231#section-4.3.2
	if c.Request.Method == http.MethodHead {
		return
	}
	status := "Service is all good"
	c.JSONP(http.StatusOK, gin.H{
		"Status": status,
	})
}

package roboapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func apiResponse(c *gin.Context, respData interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"error":   nil,
			"data":    respData,
		})
	}
}

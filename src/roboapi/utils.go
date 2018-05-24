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

// 转换请求参数
func transSolutionParams(c *gin.Context) map[string]interface{} {
	params := make(map[string]interface{})

	for key, value := range c.Request.URL.Query() {
		params[key] = value[0]
	}
	return params
}

package roboapi

import (
	"github.com/gin-gonic/gin"
)

func (e *GinEnv) rpcIndex(c *gin.Context) {
	apiResponse(c, RespSuccess, nil, "")
}

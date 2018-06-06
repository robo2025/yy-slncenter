package roboapi

import (
	"github.com/gin-gonic/gin"
	"robodb"
)

func (e *GinEnv) rpcIndex(c *gin.Context) {
	apiResponse(c, RespSuccess, nil, "")
}

func (e *GinEnv) rpcSolution(c *gin.Context) {

	rpcParams := &robodb.SolutionRPCReqParams{}
	err := c.BindJSON(rpcParams)
	if err != nil {
		apiResponse(c, RespFailed, nil, err.Error())
		return
	}

	slnDict, err := robodb.FetchSolutionRPC(e.db, rpcParams)
	if err != nil {
		apiResponse(c, RespFailed, slnDict, err.Error())
	} else {
		apiResponse(c, RespSuccess, slnDict, "")
	}
}
func (e *GinEnv) rpcSolutionDetail(c *gin.Context) {
	resp, err := robodb.FetchSolutionRPCDetail(e.db, c)
	if err != nil {
		apiResponse(c, RespFailed, resp, err.Error())
	} else {
		apiResponse(c, RespSuccess, resp, "")
	}
}

package roboapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
10000	操作成功
10001	操作失败
20001	token不存在
20002	token过期
20003	token非法
20004	登录超时
30001	无管理员权限
*/
type RespCode int

const (
	RespSuccess       RespCode = 10000
	RespFailed        RespCode = 10001
	RespTokenNotExist RespCode = 20001
	RespTokenFailed   RespCode = 20003
	RespNoAuth        RespCode = 30001
	RespNoData        RespCode = 40001
)

// API 统一回复
func apiResponse(c *gin.Context, respCode RespCode, respData interface{}, respMsg string) {
	if respCode == RespSuccess {
		if respMsg == "" {
			respMsg = "操作成功"
		}

		c.JSON(http.StatusOK, gin.H{
			"rescode": respCode,
			"data":    respData,
			"msg":     respMsg,
			"success": true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"rescode": respCode,
			"data":    nil,
			"msg":     respMsg,
			"success": false,
		})
	}
}

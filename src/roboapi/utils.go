package roboapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
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

/*
1: 普通用户
2: 供应商
3: 管理员
4: 超级管理员
*/

type UserRole int

const (
	RoleCustomer UserRole = 1
	RoleSupplier UserRole = 2
	RoleAdmin    UserRole = 3
	RoleSuper    UserRole = 4
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
		c.JSON(http.StatusBadRequest, gin.H{
			"rescode": respCode,
			"data":    nil,
			"msg":     respMsg,
			"success": false,
		})
	}
}

// 检测用户的权限和登录
func checkAuthRole(c *gin.Context, role string) error {
	isAuth := c.GetBool("isAuth")
	if !isAuth {
		err := errors.New("用户鉴权失败")
		apiResponse(c, RespTokenFailed, "", err.Error())
		return err
	}

	userRole := c.MustGet("role").(int)

	switch role {
	case "", "customer":
		if userRole != int(RoleCustomer) {
			return errors.New("用户角色不是普通用户")
		}
	case "supplier":
		if userRole != int(RoleSupplier) {
			return errors.New("用户角色不是供应商")
		}
	case "admin":
		if userRole != int(RoleAdmin) {
			return errors.New("用户角色不是管理员")
		}
	case "super":
		if userRole != int(RoleSuper) {
			return errors.New("用户角色不是超级管理员")
		}
	case "pass":
		return nil
	default:
		return errors.New("用户角色错误")
	}

	return nil
}

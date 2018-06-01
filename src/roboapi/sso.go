package roboapi

import (
	"github.com/gin-gonic/gin"
	"os"
	"roboweb"
	"encoding/json"
)

// SSO struct
type SSOJson struct {
	Data    SSOUser `json:"data"`
	Msg     string  `json:"msg"`
	Rescode string  `json:"rescode"`
}

// SSO User struct
type SSOUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	UserType int    `json:"user_type"`
}

// Gin 中间件，用户检查 Authorization
func SSOAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			apiResponse(c, RespTokenNotExist, nil, "请求中没有 Token")
			c.Set("isAuth", false)
			return
		}

		// 解析 Token
		ssoUser, err := parseTokenInfo(token)
		if err != nil {
			apiResponse(c, RespTokenFailed, nil, err.Error())
			c.Set("isAuth", false)
			return
		}

		// 设置全局参数
		c.Set("isAuth", true)
		c.Set("uid", ssoUser.ID)

		// 1:普通用户 2:供应商 3:管理员 4:超级管理员
		switch ssoUser.UserType {
		case 1:
			c.Set("user_type", "customer")
		case 2:
			c.Set("user_type", "supplier")
		case 3:
			c.Set("user_type", "admin")
		case 4:
			c.Set("user_type", "super")
		default:
			c.Set("user_type", "unknown")
		}
	}
}

// 通过 SSO 获取用户的具体信息
func parseTokenInfo(auth string) (*SSOUser, error) {
	// 定义变量
	var err error
	jsonData := make([]byte, 0)

	// 获取 SSO URL
	ssoURL := os.Getenv("SSO_URL")
	if ssoURL == "" {
		ssoURL = "https://login.robo2025.com/server/verify"
	}

	// 请求 SSO URL
	sigChan := make(chan error, 1)
	go func() {
		jsonData, err = roboweb.HttpRequest("GET", ssoURL, auth)
		sigChan <- err
	}()

	if reqErr := <-sigChan; reqErr != nil {
		return nil, reqErr
	}

	// 解析 SSO
	var ssoJSON SSOJson
	err = json.Unmarshal(jsonData, &ssoJSON)

	if err != nil {
		return nil, err
	} else {
		return &ssoJSON.Data, nil
	}
}

package roboapi

import (
	"github.com/gin-gonic/gin"
	"roboweb"
	"encoding/json"
	"os"
)

// SSO struct
type SSOJson struct {
	Data    SSOUser `json:"data"`
	Msg     string  `json:"msg"`
	Rescode string  `json:"rescode"`
}

// SSO User struct
type SSOUser struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	UserType   int    `json:"user_type"`
	IsSubuser  int    `json:"is_subuser"`
	MainUserId int    `json:"main_user_id"`
}

//SSO

// Gin 中间件，用户检查 Authorization
func SSOAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			// apiResponse(c, RespTokenNotExist, nil, "请求中没有 Token")
			c.Set("isAuth", false)
			return
		}

		// 解析 Token
		ssoUser, err := parseTokenInfo(token)
		if err != nil {
			// apiResponse(c, RespTokenFailed, nil, "用户鉴权失败")
			c.Set("isAuth", false)
			return
		}

		// 设置全局参数
		c.Set("isAuth", true)
		c.Set("uid", ssoUser.ID)
		c.Set("role", ssoUser.UserType)
		c.Set("is_subuser", ssoUser.IsSubuser)
		c.Set("main_user_id", ssoUser.MainUserId)

	}
}

// 通过 SSO 获取用户的具体信息
func parseTokenInfo(auth string) (*SSOUser, error) {
	// 定义变量
	var err error
	jsonData := make([]byte, 0)

	// 获取 SSO URL
	var ssoURL string
	if DeployMode == "production" {
		ssoURL = os.Getenv("SSO_HOST")
		if ssoURL == "" {
			ssoURL = "https://testapi.robo2025.com/sso"
		}
	} else {
		ssoURL = os.Getenv("SSO_HOST")
		if ssoURL == "" {
			ssoURL = "https://testapi.robo2025.com/sso"
		}
	}

	// 请求 SSO URL
	sigChan := make(chan error, 1)
	go func() {
		header := map[string]string{
			"Authorization": auth,
			"Content-type":  "application/json",
		}
		jsonData, err = roboweb.HttpRequest("GET", ssoURL, header, nil)
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

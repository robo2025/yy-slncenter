package roboapi

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

// Refer: https://www.jianshu.com/p/1f9915818992

// JWT 密钥
var SignKey string = "mysupersecretpassword"

// 载荷
type CustomClaims struct {
	UID int     `json:"uid"`
	SID string  `json:"sid"`
	IAT float64 `json:"iat"`
}

func parseToken(tokenString string) (*CustomClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("非法的签名算法: %v", token.Header["alg"])
		}

		// 设置签名密钥
		return []byte(SignKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		resp := &CustomClaims{
			UID: int(claims["uid"].(float64)),
			SID: claims["sid"].(string),
			IAT: claims["iat"].(float64),
		}
		return resp, nil
	}

	return nil, fmt.Errorf("无效的 Token")
}

// 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			apiResponse(c, nil, fmt.Errorf("请求中没有 Token"))
			c.Set("isAuth", false)
			return
		}

		// 解析 Token
		claims, err := parseToken(token)
		if err != nil {
			apiResponse(c, nil, err)
			c.Set("isAuth", false)
			return
		}

		c.Set("isAuth", true)
		c.Set("uid", claims.UID)
	}
}

package MiddleWare

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.JSON(200, gin.H{"data": "home"})
				c.Next()
				return
			}
		}
		//返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		//若验证不通过，不在调用后续的函数处理
		c.Abort()
		return
	}
}

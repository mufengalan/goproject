package MiddleWare

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		if cookie, err := context.Cookie("abc"); err == nil {
			if cookie == "123" {
				context.Next()
				return
			}
		}
		//返回错误
		context.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		//若验证不通过，不在调用后续的函数处理
		context.Abort()
		return
	}
}

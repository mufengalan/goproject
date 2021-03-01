package main

import (
	"github.com/gin-gonic/gin"
	"project/App/Login"
)

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// JSON绑定
	r.POST("/loginForm", Login.AppLogin)
	r.Run(":8000")
}
